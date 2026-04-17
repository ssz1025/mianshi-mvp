package dto

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"omitempty"`
	Avatar   string `json:"avatar" binding:"omitempty"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone" binding:"omitempty"`
	OpenID   string `json:"openid" binding:"omitempty"`
}

type UpdateUserRequest struct {
	ID            int64   `uri:"id" binding:"required"`
	Username      *string `json:"username" binding:"omitempty"`
	Nickname      *string `json:"nickname" binding:"omitempty"`
	Avatar        *string `json:"avatar" binding:"omitempty"`
	Phone         *string `json:"phone" binding:"omitempty"`
	OpenID        *string `json:"openid" binding:"omitempty"`
	IsVIP         *bool   `json:"is_vip" binding:"omitempty"`
	VIPExpireTime *string `json:"vip_expire_time" binding:"omitempty"`
	Integral      *int    `json:"integral" binding:"omitempty,gte=0"`
}

// GetUserRequest 获取用户请求
type GetUserRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

// DeleteUserRequest 删除用户请求
type DeleteUserRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

// ListUsersRequest 用户列表请求
type ListUsersRequest struct {
	Page     int `form:"page" binding:"omitempty,gte=1"`
	PageSize int `form:"page_size" binding:"omitempty,gte=1,lte=100"`
}

// SetDefaults 设置默认值
func (r *ListUsersRequest) SetDefaults() {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.PageSize < 1 || r.PageSize > 100 {
		r.PageSize = 10
	}
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token    string `json:"token"`
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}

// UserResponse 用户响应
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,min=6"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type UpdateProfileRequest struct {
	Nickname *string `json:"nickname" binding:"omitempty"`
	Avatar   *string `json:"avatar" binding:"omitempty"`
	Phone    *string `json:"phone" binding:"omitempty"`
}

type UserStatsResponse struct {
	TotalQuestions  int `json:"total_questions"`
	MasterQuestions int `json:"master_questions"`
	FavoriteCount   int `json:"favorite_count"`
	SearchCount     int `json:"search_count"`
}

type UserResponse struct {
	ID            int64  `json:"id"`
	Username      string `json:"username"`
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	Phone         string `json:"phone"`
	OpenID        string `json:"openid"`
	IsVIP         bool   `json:"is_vip"`
	VIPExpireTime string `json:"vip_expire_time"`
	Integral      int    `json:"integral"`
	IsDeleted     bool   `json:"is_deleted"`
	CreateTime    string `json:"create_time"`
	UpdateTime    string `json:"update_time"`
}
