package model

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

type PracticeRoute struct {
	ID              int64        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string       `json:"name" gorm:"type:varchar(100)"`
	Title           string       `json:"title" gorm:"type:varchar(100)"`
	Icon            string       `json:"icon" gorm:"type:varchar(255)"`
	Color           string       `json:"color" gorm:"type:varchar(100)"`
	Description     string       `json:"description" gorm:"type:text"`
	TargetLevel     string       `json:"target_level" gorm:"type:varchar(100)"`
	SuitableFor     StringArray  `json:"suitable_for" gorm:"type:text"`
	Skills          StringArray  `json:"skills" gorm:"type:text"`
	InterviewWeight string       `json:"interview_weight" gorm:"type:varchar(50)"`
	Sort            int          `json:"sort" gorm:"default:0"`
	IsDeleted       bool         `json:"is_deleted" gorm:"default:false"`
	CreateTime      time.Time    `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime      time.Time    `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
	Phases          []RoutePhase `json:"phases" gorm:"foreignKey:RouteID;references:ID"`
}

type RoutePhase struct {
	ID          int64       `json:"id" gorm:"primaryKey;autoIncrement"`
	RouteID     int64       `json:"route_id" gorm:"index"`
	Phase       string      `json:"phase" gorm:"type:varchar(100)"`
	Duration    string      `json:"duration" gorm:"type:varchar(50)"`
	Topics      StringArray `json:"topics" gorm:"type:text"`
	Description string      `json:"description" gorm:"type:text"`
	Sort        int         `json:"sort" gorm:"default:0"`
	CreateTime  time.Time   `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime  time.Time   `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
}

type StringArray []string

func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = []string{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return s.scanString(string(v))
	case string:
		return s.scanString(v)
	case []string:
		*s = v
		return nil
	}

	return nil
}

func (s *StringArray) scanString(val string) error {
	if val == "" || val == "{}" || val == "[]" {
		*s = []string{}
		return nil
	}

	if val[0] == '[' {
		var result []string
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			*s = result
			return nil
		}
	}

	val = strings.TrimPrefix(val, "{")
	val = strings.TrimSuffix(val, "}")

	if val == "" {
		*s = []string{}
		return nil
	}

	parts := strings.Split(val, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.Trim(p, "\"")
		if p != "" {
			result = append(result, p)
		}
	}
	*s = result
	return nil
}

func (s StringArray) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "[]", nil
	}
	data, err := json.Marshal([]string(s))
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (PracticeRoute) TableName() string {
	return "practice_route"
}

func (RoutePhase) TableName() string {
	return "route_phase"
}
