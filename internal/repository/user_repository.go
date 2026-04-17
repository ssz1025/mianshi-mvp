package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/d60-Lab/gin-template/internal/model"
)

// UserRepository 用户仓储接口
type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, offset, limit int) ([]*model.User, error)
	CountQuestionRecords(ctx context.Context, userID int64) (int, error)
	CountMasterRecords(ctx context.Context, userID int64) (int, error)
	CountFavorites(ctx context.Context, userID int64) (int, error)
	CountSearchHistory(ctx context.Context, userID int64) (int, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建用户仓储实例
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, "id = ?", id).Error
}

func (r *userRepository) List(ctx context.Context, offset, limit int) ([]*model.User, error) {
	var users []*model.User
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func (r *userRepository) CountQuestionRecords(ctx context.Context, userID int64) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Table("user_question_record").Where("user_id = ?", userID).Count(&count).Error
	return int(count), err
}

func (r *userRepository) CountMasterRecords(ctx context.Context, userID int64) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Table("user_question_record").Where("user_id = ? AND is_master = ?", userID, true).Count(&count).Error
	return int(count), err
}

func (r *userRepository) CountFavorites(ctx context.Context, userID int64) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Table("user_favorite").Where("user_id = ?", userID).Count(&count).Error
	return int(count), err
}

func (r *userRepository) CountSearchHistory(ctx context.Context, userID int64) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Table("user_search_history").Where("user_id = ?", userID).Count(&count).Error
	return int(count), err
}
