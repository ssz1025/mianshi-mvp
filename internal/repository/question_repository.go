package repository

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/d60-Lab/gin-template/internal/model"
)

type QuestionRepository interface {
	ListBanks(ctx context.Context, category, search string, offset, limit int) ([]*model.QuestionBank, int64, error)
	GetBankByID(ctx context.Context, id int64) (*model.QuestionBank, error)
	ListBankQuestions(ctx context.Context, bankID int64, difficulty int, search string, offset, limit int) ([]*model.Question, int64, error)
	GetQuestionByID(ctx context.Context, id int64) (*model.Question, error)
	ListHotQuestions(ctx context.Context, category string, difficulty int, tag string, sortBy string, offset, limit int) ([]*model.Question, int64, error)
	GetRelatedQuestions(ctx context.Context, bankID, questionID int64, limit int) ([]*model.Question, error)
	GetQuestionTags(ctx context.Context, questionID int64) ([]string, error)
	GetQuestionCategory(ctx context.Context, questionID int64) (string, error)
	GetQuestionBankIDs(ctx context.Context, questionID int64) ([]int64, error)
	ListUserRecords(ctx context.Context, userID int64, filter string, offset, limit int) ([]*model.UserQuestionRecord, int64, error)
	CreateRecord(ctx context.Context, userID, questionID int64) (*model.UserQuestionRecord, error)
	ToggleMaster(ctx context.Context, userID, questionID int64, isMaster bool) error
	GetRecordByQuestion(ctx context.Context, userID, questionID int64) (*model.UserQuestionRecord, error)
	ToggleFavorite(ctx context.Context, userID, questionID int64) (bool, error)
	ListFavorites(ctx context.Context, userID int64, offset, limit int) ([]*model.UserFavorite, int64, error)
	CheckFavorite(ctx context.Context, userID, questionID int64) (bool, error)
}

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) ListBanks(ctx context.Context, category, search string, offset, limit int) ([]*model.QuestionBank, int64, error) {
	var banks []*model.QuestionBank
	var total int64

	query := r.db.WithContext(ctx).Where("is_deleted = ?", false)

	if category != "" && category != "all" {
		query = query.Where("category = ?", category)
	}
	if search != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Model(&model.QuestionBank{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("sort ASC, id ASC").Offset(offset).Limit(limit).Find(&banks).Error
	return banks, total, err
}

func (r *questionRepository) GetBankByID(ctx context.Context, id int64) (*model.QuestionBank, error) {
	var bank model.QuestionBank
	err := r.db.WithContext(ctx).Where("id = ? AND is_deleted = ?", id, false).First(&bank).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &bank, nil
}

func (r *questionRepository) ListBankQuestions(ctx context.Context, bankID int64, difficulty int, search string, offset, limit int) ([]*model.Question, int64, error) {
	var questions []*model.Question
	var total int64

	subQuery := r.db.WithContext(ctx).
		Model(&model.QuestionBankQuestion{}).
		Select("question_id").
		Where("bank_id = ?", bankID)

	query := r.db.WithContext(ctx).
		Where("id IN (?)", subQuery).
		Where("is_deleted = ?", false)

	if difficulty > 0 {
		query = query.Where("difficulty = ?", difficulty)
	}
	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	if err := query.Model(&model.Question{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("id ASC").Offset(offset).Limit(limit).Find(&questions).Error
	return questions, total, err
}

func (r *questionRepository) GetQuestionByID(ctx context.Context, id int64) (*model.Question, error) {
	var question model.Question
	err := r.db.WithContext(ctx).Where("id = ? AND is_deleted = ?", id, false).First(&question).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &question, nil
}

func (r *questionRepository) ListHotQuestions(ctx context.Context, category string, difficulty int, tag string, sortBy string, offset, limit int) ([]*model.Question, int64, error) {
	var questions []*model.Question
	var total int64

	query := r.db.WithContext(ctx).Where("question.is_deleted = ?", false)

	if category != "" || tag != "" {
		tagSubQuery := r.db.Model(&model.QuestionTagRelation{}).
			Select("question_id").
			Joins("JOIN tag ON tag.id = question_tag.tag_id")

		if category != "" {
			tagSubQuery = tagSubQuery.Where("tag.category = ?", category)
		}
		if tag != "" {
			tagSubQuery = tagSubQuery.Where("tag.name = ?", tag)
		}
		query = query.Where("question.id IN (?)", tagSubQuery)
	}

	if difficulty > 0 {
		query = query.Where("question.difficulty = ?", difficulty)
	}

	if err := query.Model(&model.Question{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	orderClause := "view_count DESC, id ASC"
	switch sortBy {
	case "newest":
		orderClause = "id DESC"
	case "difficulty":
		orderClause = "difficulty DESC, id ASC"
	case "star":
		orderClause = "star_count DESC, id ASC"
	}

	err := query.Order(orderClause).Offset(offset).Limit(limit).Find(&questions).Error
	return questions, total, err
}

func (r *questionRepository) GetRelatedQuestions(ctx context.Context, bankID, questionID int64, limit int) ([]*model.Question, error) {
	var questions []*model.Question

	subQuery := r.db.WithContext(ctx).
		Model(&model.QuestionBankQuestion{}).
		Select("question_id").
		Where("bank_id = ? AND question_id != ?", bankID, questionID)

	err := r.db.WithContext(ctx).
		Where("id IN (?) AND is_deleted = ?", subQuery, false).
		Order("view_count DESC").
		Limit(limit).
		Find(&questions).Error
	return questions, err
}

func (r *questionRepository) GetQuestionTags(ctx context.Context, questionID int64) ([]string, error) {
	var tagNames []string
	err := r.db.WithContext(ctx).
		Model(&model.QuestionTag{}).
		Joins("JOIN question_tag AS qt ON qt.tag_id = tag.id").
		Where("qt.question_id = ?", questionID).
		Pluck("tag.name", &tagNames).Error
	return tagNames, err
}

func (r *questionRepository) GetQuestionCategory(ctx context.Context, questionID int64) (string, error) {
	var category string
	err := r.db.WithContext(ctx).
		Model(&model.QuestionTag{}).
		Joins("JOIN question_tag AS qt ON qt.tag_id = tag.id").
		Where("qt.question_id = ?", questionID).
		Limit(1).
		Pluck("tag.category", &category).Error
	return category, err
}

func (r *questionRepository) GetQuestionBankIDs(ctx context.Context, questionID int64) ([]int64, error) {
	var bankIDs []int64
	err := r.db.WithContext(ctx).
		Model(&model.QuestionBankQuestion{}).
		Where("question_id = ?", questionID).
		Pluck("bank_id", &bankIDs).Error
	return bankIDs, err
}

func (r *questionRepository) ListUserRecords(ctx context.Context, userID int64, filter string, offset, limit int) ([]*model.UserQuestionRecord, int64, error) {
	var records []*model.UserQuestionRecord
	var total int64

	query := r.db.WithContext(ctx).Where("user_id = ?", userID)

	switch filter {
	case "master":
		query = query.Where("is_master = ?", true)
	case "not-master":
		query = query.Where("is_master = ?", false)
	}

	if err := query.Model(&model.UserQuestionRecord{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("last_view_time DESC").Offset(offset).Limit(limit).Find(&records).Error
	return records, total, err
}

func (r *questionRepository) CreateRecord(ctx context.Context, userID, questionID int64) (*model.UserQuestionRecord, error) {
	var record model.UserQuestionRecord
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND question_id = ?", userID, questionID).
		First(&record).Error

	if err == gorm.ErrRecordNotFound {
		record = model.UserQuestionRecord{
			UserID:       userID,
			QuestionID:   questionID,
			LastViewTime: time.Now(),
		}
		err = r.db.WithContext(ctx).Create(&record).Error
		return &record, err
	}

	if err != nil {
		return nil, err
	}

	record.LastViewTime = time.Now()
	err = r.db.WithContext(ctx).Save(&record).Error
	return &record, err
}

func (r *questionRepository) ToggleMaster(ctx context.Context, userID, questionID int64, isMaster bool) error {
	return r.db.WithContext(ctx).
		Model(&model.UserQuestionRecord{}).
		Where("user_id = ? AND question_id = ?", userID, questionID).
		Update("is_master", isMaster).Error
}

func (r *questionRepository) GetRecordByQuestion(ctx context.Context, userID, questionID int64) (*model.UserQuestionRecord, error) {
	var record model.UserQuestionRecord
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND question_id = ?", userID, questionID).
		First(&record).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &record, err
}

func (r *questionRepository) ToggleFavorite(ctx context.Context, userID, questionID int64) (bool, error) {
	var fav model.UserFavorite
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND question_id = ?", userID, questionID).
		First(&fav).Error

	if err == gorm.ErrRecordNotFound {
		newFav := model.UserFavorite{
			UserID:     userID,
			QuestionID: questionID,
		}
		if err := r.db.WithContext(ctx).Create(&newFav).Error; err != nil {
			return false, err
		}
		return true, nil
	}

	if err != nil {
		return false, err
	}

	if err := r.db.WithContext(ctx).Delete(&fav).Error; err != nil {
		return false, err
	}
	return false, nil
}

func (r *questionRepository) ListFavorites(ctx context.Context, userID int64, offset, limit int) ([]*model.UserFavorite, int64, error) {
	var favorites []*model.UserFavorite
	var total int64

	query := r.db.WithContext(ctx).Where("user_id = ?", userID)

	if err := query.Model(&model.UserFavorite{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("create_time DESC").Offset(offset).Limit(limit).Find(&favorites).Error
	return favorites, total, err
}

func (r *questionRepository) CheckFavorite(ctx context.Context, userID, questionID int64) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.UserFavorite{}).
		Where("user_id = ? AND question_id = ?", userID, questionID).
		Count(&count).Error
	return count > 0, err
}
