package model

import (
	"time"
)

type PracticeRoute struct {
	ID              int64         `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string        `json:"name" gorm:"type:varchar(100)"`
	Title           string        `json:"title" gorm:"type:varchar(100)"`
	Icon            string        `json:"icon" gorm:"type:varchar(255)"`
	Color           string        `json:"color" gorm:"type:varchar(100)"`
	Description     string        `json:"description" gorm:"type:text"`
	TargetLevel     string        `json:"target_level" gorm:"type:varchar(100)"`
	SuitableFor     StringArray   `json:"suitable_for" gorm:"type:varchar(255)[]"`
	Skills          StringArray   `json:"skills" gorm:"type:varchar(255)[]"`
	InterviewWeight string        `json:"interview_weight" gorm:"type:varchar(50)"`
	Sort            int           `json:"sort" gorm:"default:0"`
	IsDeleted       bool          `json:"is_deleted" gorm:"default:false"`
	CreateTime      time.Time     `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime      time.Time     `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	Phases          []RoutePhase  `json:"phases" gorm:"foreignKey:RouteID;references:ID"`
}

type RoutePhase struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	RouteID     int64     `json:"route_id" gorm:"index"`
	Phase       string    `json:"phase" gorm:"type:varchar(100)"`
	Duration    string    `json:"duration" gorm:"type:varchar(50)"`
	Topics      StringArray `json:"topics" gorm:"type:varchar(255)[]"`
	Description string    `json:"description" gorm:"type:text"`
	Sort        int       `json:"sort" gorm:"default:0"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
}

type StringArray []string

func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = []string{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	if len(bytes) == 0 {
		*s = []string{}
		return nil
	}
	result := string(bytes)
	if result == "{}" || result == "[]" {
		*s = []string{}
		return nil
	}
	*s = []string{result}
	return nil
}

func (s StringArray) Value() (interface{}, error) {
	if len(s) == 0 {
		return "{}", nil
	}
	return s, nil
}

func (PracticeRoute) TableName() string {
	return "practice_route"
}

func (RoutePhase) TableName() string {
	return "route_phase"
}
