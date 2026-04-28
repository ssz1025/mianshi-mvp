package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/d60-Lab/gin-template/internal/dto"
	"github.com/d60-Lab/gin-template/internal/model"
	"github.com/d60-Lab/gin-template/internal/repository"
)

var (
	ErrBankNotFound     = errors.New("bank not found")
	ErrQuestionNotFound = errors.New("question not found")
)

func difficultyToInt(d string) int {
	switch d {
	case "简单":
		return 1
	case "中等":
		return 2
	case "困难":
		return 3
	default:
		if n, err := strconv.Atoi(d); err == nil {
			return n
		}
		return 0
	}
}

func difficultyToStr(d int) string {
	switch d {
	case 1:
		return "简单"
	case 2:
		return "中等"
	case 3:
		return "困难"
	default:
		return "中等"
	}
}

type QuestionService interface {
	ListBanks(ctx context.Context, req *dto.ListBanksRequest) (*dto.PaginatedResponse, error)
	GetBankByID(ctx context.Context, id int64) (*dto.QuestionBankResponse, error)
	ListBankQuestions(ctx context.Context, req *dto.ListBankQuestionsRequest) (*dto.PaginatedResponse, error)
	GetQuestionByID(ctx context.Context, id int64) (*dto.QuestionDetail, error)
	ListHotQuestions(ctx context.Context, req *dto.HotQuestionsRequest) (*dto.PaginatedResponse, error)
	ListQuestions(ctx context.Context, req *dto.ListQuestionsRequest) (*dto.PaginatedResponse, error)
	ListUserRecords(ctx context.Context, userID int64, req *dto.ListRecordsRequest) (*dto.PaginatedResponse, error)
	CreateRecord(ctx context.Context, userID int64, req *dto.CreateRecordRequest) (*dto.QuestionRecordItem, error)
	ToggleMaster(ctx context.Context, userID int64, req *dto.ToggleMasterRequest) (*dto.QuestionRecordItem, error)
	ToggleFavorite(ctx context.Context, userID int64, req *dto.ToggleFavoriteRequest) (bool, error)
	ListFavorites(ctx context.Context, userID int64, req *dto.ListRecordsRequest) (*dto.PaginatedResponse, error)
}

type questionService struct {
	questionRepo repository.QuestionRepository
}

func NewQuestionService(questionRepo repository.QuestionRepository) QuestionService {
	return &questionService{
		questionRepo: questionRepo,
	}
}

func (s *questionService) ListBanks(ctx context.Context, req *dto.ListBanksRequest) (*dto.PaginatedResponse, error) {
	req.SetDefaults()
	offset := (req.Page - 1) * req.PageSize

	banks, total, err := s.questionRepo.ListBanks(ctx, req.Category, req.Search, offset, req.PageSize)
	if err != nil {
		return nil, err
	}

	list := make([]*dto.QuestionBankResponse, len(banks))
	for i, bank := range banks {
		list[i] = toBankResponse(bank)
	}

	return &dto.PaginatedResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *questionService) GetBankByID(ctx context.Context, id int64) (*dto.QuestionBankResponse, error) {
	bank, err := s.questionRepo.GetBankByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if bank == nil {
		return nil, ErrBankNotFound
	}
	return toBankResponse(bank), nil
}

func (s *questionService) ListBankQuestions(ctx context.Context, req *dto.ListBankQuestionsRequest) (*dto.PaginatedResponse, error) {
	req.SetDefaults()
	offset := (req.Page - 1) * req.PageSize

	difficulty := difficultyToInt(req.Difficulty)
	questions, total, err := s.questionRepo.ListBankQuestions(ctx, req.ID, difficulty, req.Search, offset, req.PageSize)
	if err != nil {
		return nil, err
	}

	list := make([]*dto.QuestionListItem, len(questions))
	for i, q := range questions {
		item := toQuestionListItem(q)
		tags, _ := s.questionRepo.GetQuestionTags(ctx, q.ID)
		item.Tags = tags
		cat, _ := s.questionRepo.GetQuestionCategory(ctx, q.ID)
		item.Category = cat
		list[i] = item
	}

	return &dto.PaginatedResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *questionService) GetQuestionByID(ctx context.Context, id int64) (*dto.QuestionDetail, error) {
	question, err := s.questionRepo.GetQuestionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, ErrQuestionNotFound
	}

	tags, _ := s.questionRepo.GetQuestionTags(ctx, question.ID)
	category, _ := s.questionRepo.GetQuestionCategory(ctx, question.ID)

	related := make([]dto.RelatedQuestion, 0)
	bankQuestions, _ := s.questionRepo.GetRelatedQuestions(ctx, 0, question.ID, 3)
	if bankQuestions != nil {
		for _, rq := range bankQuestions {
			related = append(related, dto.RelatedQuestion{
				ID:    rq.ID,
				Title: rq.Title,
			})
		}
	}

	return &dto.QuestionDetail{
		ID:               question.ID,
		Title:            question.Title,
		Content:          question.Content,
		Category:         category,
		Difficulty:       difficultyToStr(question.Difficulty),
		Tags:             tags,
		IsVIP:            question.IsVIP,
		ViewCount:        question.ViewCount,
		StarCount:        question.StarCount,
		LikeCount:        question.LikeCount,
		Answer:           question.Answer,
		Explanation:      question.Explanation,
		RelatedQuestions: related,
	}, nil
}

func (s *questionService) ListHotQuestions(ctx context.Context, req *dto.HotQuestionsRequest) (*dto.PaginatedResponse, error) {
	req.SetDefaults()
	offset := (req.Page - 1) * req.PageSize

	difficulty := difficultyToInt(req.Difficulty)
	questions, total, err := s.questionRepo.ListHotQuestions(ctx, req.Category, difficulty, req.Tag, req.SortBy, offset, req.PageSize)
	if err != nil {
		return nil, err
	}

	list := make([]*dto.QuestionListItem, len(questions))
	for i, q := range questions {
		item := toQuestionListItem(q)
		tags, _ := s.questionRepo.GetQuestionTags(ctx, q.ID)
		item.Tags = tags
		cat, _ := s.questionRepo.GetQuestionCategory(ctx, q.ID)
		item.Category = cat
		list[i] = item
	}

	return &dto.PaginatedResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *questionService) ListQuestions(ctx context.Context, req *dto.ListQuestionsRequest) (*dto.PaginatedResponse, error) {
	req.SetDefaults()
	offset := (req.Page - 1) * req.PageSize

	difficulty := difficultyToInt(req.Difficulty)

	if req.BankID > 0 {
		questions, total, err := s.questionRepo.ListBankQuestions(ctx, req.BankID, difficulty, req.Search, offset, req.PageSize)
		if err != nil {
			return nil, err
		}
		list := make([]*dto.QuestionListItem, len(questions))
		for i, q := range questions {
			item := toQuestionListItem(q)
			tags, _ := s.questionRepo.GetQuestionTags(ctx, q.ID)
			item.Tags = tags
			cat, _ := s.questionRepo.GetQuestionCategory(ctx, q.ID)
			item.Category = cat
			list[i] = item
		}
		return &dto.PaginatedResponse{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, nil
	}

	questions, total, err := s.questionRepo.ListHotQuestions(ctx, req.Category, difficulty, req.Tag, req.SortBy, offset, req.PageSize)
	if err != nil {
		return nil, err
	}

	list := make([]*dto.QuestionListItem, len(questions))
	for i, q := range questions {
		item := toQuestionListItem(q)
		tags, _ := s.questionRepo.GetQuestionTags(ctx, q.ID)
		item.Tags = tags
		cat, _ := s.questionRepo.GetQuestionCategory(ctx, q.ID)
		item.Category = cat
		list[i] = item
	}

	return &dto.PaginatedResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func toBankResponse(bank *model.QuestionBank) *dto.QuestionBankResponse {
	updateTime := ""
	if !bank.UpdateTime.IsZero() {
		updateTime = bank.UpdateTime.Format(time.DateOnly)
	}
	return &dto.QuestionBankResponse{
		ID:             bank.ID,
		Title:          bank.Title,
		Description:    bank.Description,
		Icon:           bank.Icon,
		Category:       bank.Category,
		IsVIP:          bank.IsVIP,
		TotalQuestions: bank.TotalQuestions,
		UpdateTime:     updateTime,
	}
}

func toQuestionListItem(q *model.Question) *dto.QuestionListItem {
	return &dto.QuestionListItem{
		ID:         q.ID,
		Title:      q.Title,
		Difficulty: difficultyToStr(q.Difficulty),
		IsVIP:      q.IsVIP,
		ViewCount:  q.ViewCount,
		StarCount:  q.StarCount,
		LikeCount:  q.LikeCount,
	}
}

func (s *questionService) ListUserRecords(ctx context.Context, userID int64, req *dto.ListRecordsRequest) (*dto.PaginatedResponse, error) {
	req.SetDefaults()
	offset := (req.Page - 1) * req.PageSize

	records, total, err := s.questionRepo.ListUserRecords(ctx, userID, req.Filter, offset, req.PageSize)
	if err != nil {
		return nil, err
	}

	list := make([]*dto.QuestionRecordItem, len(records))
	for i, record := range records {
		tags, _ := s.questionRepo.GetQuestionTags(ctx, record.QuestionID)
		category, _ := s.questionRepo.GetQuestionCategory(ctx, record.QuestionID)
		question, _ := s.questionRepo.GetQuestionByID(ctx, record.QuestionID)

		title := ""
		difficulty := ""
		if question != nil {
			title = question.Title
			difficulty = difficultyToStr(question.Difficulty)
		}

		list[i] = &dto.QuestionRecordItem{
			ID:                 record.ID,
			QuestionID:         record.QuestionID,
			QuestionTitle:      title,
			QuestionDifficulty: difficulty,
			QuestionCategory:   category,
			QuestionTags:       tags,
			IsMaster:           record.IsMaster,
			LastViewTime:       formatTime(record.LastViewTime),
			CreateTime:         formatTime(record.CreateTime),
		}
	}

	return &dto.PaginatedResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}

func (s *questionService) CreateRecord(ctx context.Context, userID int64, req *dto.CreateRecordRequest) (*dto.QuestionRecordItem, error) {
	record, err := s.questionRepo.CreateRecord(ctx, userID, req.QuestionID)
	if err != nil {
		return nil, err
	}

	tags, _ := s.questionRepo.GetQuestionTags(ctx, record.QuestionID)
	category, _ := s.questionRepo.GetQuestionCategory(ctx, record.QuestionID)
	question, _ := s.questionRepo.GetQuestionByID(ctx, record.QuestionID)

	title := ""
	difficulty := ""
	if question != nil {
		title = question.Title
		difficulty = difficultyToStr(question.Difficulty)
	}

	return &dto.QuestionRecordItem{
		ID:                 record.ID,
		QuestionID:         record.QuestionID,
		QuestionTitle:      title,
		QuestionDifficulty: difficulty,
		QuestionCategory:   category,
		QuestionTags:       tags,
		IsMaster:           record.IsMaster,
		LastViewTime:       formatTime(record.LastViewTime),
		CreateTime:         formatTime(record.CreateTime),
	}, nil
}

func (s *questionService) ToggleMaster(ctx context.Context, userID int64, req *dto.ToggleMasterRequest) (*dto.QuestionRecordItem, error) {
	err := s.questionRepo.ToggleMaster(ctx, userID, req.QuestionID, req.IsMaster)
	if err != nil {
		return nil, err
	}

	record, err := s.questionRepo.GetRecordByQuestion(ctx, userID, req.QuestionID)
	if err != nil || record == nil {
		return nil, err
	}

	tags, _ := s.questionRepo.GetQuestionTags(ctx, record.QuestionID)
	category, _ := s.questionRepo.GetQuestionCategory(ctx, record.QuestionID)
	question, _ := s.questionRepo.GetQuestionByID(ctx, record.QuestionID)

	title := ""
	difficulty := ""
	if question != nil {
		title = question.Title
		difficulty = difficultyToStr(question.Difficulty)
	}

	return &dto.QuestionRecordItem{
		ID:                 record.ID,
		QuestionID:         record.QuestionID,
		QuestionTitle:      title,
		QuestionDifficulty: difficulty,
		QuestionCategory:   category,
		QuestionTags:       tags,
		IsMaster:           record.IsMaster,
		LastViewTime:       formatTime(record.LastViewTime),
		CreateTime:         formatTime(record.CreateTime),
	}, nil
}

func (s *questionService) ToggleFavorite(ctx context.Context, userID int64, req *dto.ToggleFavoriteRequest) (bool, error) {
	return false, nil
}

func (s *questionService) ListFavorites(ctx context.Context, userID int64, req *dto.ListRecordsRequest) (*dto.PaginatedResponse, error) {
	return &dto.PaginatedResponse{
		List:     []*dto.FavoriteItem{},
		Total:    0,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
