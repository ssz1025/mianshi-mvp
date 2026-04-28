package model

import (
	"time"
)

// User 用户模型
type User struct {
	ID            int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Username      string    `json:"username" gorm:"uniqueIndex:idx_user_username;type:varchar(50)"`
	Nickname      string    `json:"nickname" gorm:"type:varchar(50)"`
	Avatar        string    `json:"avatar" gorm:"type:varchar(255)"`
	Password      string    `json:"-" gorm:"type:varchar(100)"`
	Phone         *string   `json:"phone" gorm:"uniqueIndex:idx_user_phone;type:varchar(20)"`
	OpenID        *string   `json:"openid" gorm:"column:openid;uniqueIndex:idx_user_openid;type:varchar(100)"`
	IsVIP         bool      `json:"is_vip" gorm:"column:is_vip;default:false"`
	VIPExpireTime time.Time `json:"vip_expire_time" gorm:"column:vip_expire_time"`
	Integral      int       `json:"integral" gorm:"default:0"`
	IsDeleted     bool      `json:"is_deleted" gorm:"default:false"`
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime    time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}
