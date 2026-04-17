package service

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/d60-Lab/gin-template/internal/dto"
	"github.com/d60-Lab/gin-template/internal/model"
	"github.com/d60-Lab/gin-template/internal/repository"
	"github.com/d60-Lab/gin-template/pkg/config"
	"github.com/d60-Lab/gin-template/pkg/jwt"
)

var (
	ErrUserExists      = errors.New("user already exists")
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
)

// UserService 用户服务接口
type UserService interface {
	Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error)
	GetByID(ctx context.Context, id int64) (*dto.UserResponse, error)
	Update(ctx context.Context, id int64, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
	Delete(ctx context.Context, id int64) error
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
	List(ctx context.Context, page, pageSize int) ([]*dto.UserResponse, error)
	ChangePassword(ctx context.Context, userID int64, req *dto.ChangePasswordRequest) error
	GetUserStats(ctx context.Context, userID int64) (*dto.UserStatsResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
	cfg      *config.Config
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo repository.UserRepository, cfg *config.Config) UserService {
	return &userService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *userService) Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	// 检查用户名是否已存在
	existingUser, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrUserExists
	}

	// 加密密码
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := &model.User{
		Username: req.Username,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Password: hashedPassword,
	}
	if req.Phone != "" {
		user.Phone = &req.Phone
	}
	if req.OpenID != "" {
		user.OpenID = &req.OpenID
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return s.toUserResponse(user), nil
}

func (s *userService) GetByID(ctx context.Context, id int64) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return s.toUserResponse(user), nil
}

func (s *userService) Update(ctx context.Context, id int64, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	// 更新字段
	if req.Username != nil {
		user.Username = *req.Username
	}
	if req.Nickname != nil {
		user.Nickname = *req.Nickname
	}
	if req.Avatar != nil {
		user.Avatar = *req.Avatar
	}
	if req.Phone != nil {
		user.Phone = req.Phone
	}
	if req.OpenID != nil {
		user.OpenID = req.OpenID
	}
	if req.IsVIP != nil {
		user.IsVIP = *req.IsVIP
	}
	if req.VIPExpireTime != nil {
		expireTime, err := time.Parse(time.RFC3339, *req.VIPExpireTime)
		if err == nil {
			user.VIPExpireTime = expireTime
		}
	}
	if req.Integral != nil {
		user.Integral = *req.Integral
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return s.toUserResponse(user), nil
}

func (s *userService) Delete(ctx context.Context, id int64) error {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	return s.userRepo.Delete(ctx, id)
}

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	// 验证密码
	if !checkPasswordHash(req.Password, user.Password) {
		return nil, ErrInvalidPassword
	}

	// 生成 JWT token
	token, err := jwt.GenerateToken(user.ID, user.Username, s.cfg.JWT.Secret, s.cfg.JWT.Expire) // pragma: allowlist secret
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token:    token,
		UserID:   user.ID,
		Username: user.Username,
	}, nil
}

func (s *userService) List(ctx context.Context, page, pageSize int) ([]*dto.UserResponse, error) {
	offset := (page - 1) * pageSize
	users, err := s.userRepo.List(ctx, offset, pageSize)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.UserResponse, len(users))
	for i, user := range users {
		responses[i] = s.toUserResponse(user)
	}

	return responses, nil
}

func (s *userService) ChangePassword(ctx context.Context, userID int64, req *dto.ChangePasswordRequest) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	if !checkPasswordHash(req.OldPassword, user.Password) {
		return ErrInvalidPassword
	}

	hashedPassword, err := hashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.userRepo.Update(ctx, user)
}

func (s *userService) GetUserStats(ctx context.Context, userID int64) (*dto.UserStatsResponse, error) {
	totalQuestions, err := s.userRepo.CountQuestionRecords(ctx, userID)
	if err != nil {
		return nil, err
	}
	masterQuestions, err := s.userRepo.CountMasterRecords(ctx, userID)
	if err != nil {
		return nil, err
	}
	favoriteCount, err := s.userRepo.CountFavorites(ctx, userID)
	if err != nil {
		return nil, err
	}
	searchCount, err := s.userRepo.CountSearchHistory(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &dto.UserStatsResponse{
		TotalQuestions:  totalQuestions,
		MasterQuestions: masterQuestions,
		FavoriteCount:   favoriteCount,
		SearchCount:     searchCount,
	}, nil
}

func (s *userService) toUserResponse(user *model.User) *dto.UserResponse {
	var phone string
	if user.Phone != nil {
		phone = *user.Phone
	}
	var openID string
	if user.OpenID != nil {
		openID = *user.OpenID
	}
	return &dto.UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		Nickname:      user.Nickname,
		Avatar:        user.Avatar,
		Phone:         phone,
		OpenID:       openID,
		IsVIP:         user.IsVIP,
		VIPExpireTime: user.VIPExpireTime.Format("2006-01-02 15:04:05"),
		Integral:      user.Integral,
		IsDeleted:     user.IsDeleted,
		CreateTime:    user.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:    user.UpdateTime.Format("2006-01-02 15:04:05"),
	}
}

// hashPassword 加密密码
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash 验证密码
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
