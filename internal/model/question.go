package model

import "time"

type QuestionBank struct {
	ID             int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title          string    `json:"title" gorm:"column:name;type:varchar(100);not null"`
	Description    string    `json:"description" gorm:"type:text"`
	Icon           string    `json:"icon" gorm:"type:varchar(255)"`
	Category       string    `json:"category" gorm:"type:varchar(50);index"`
	IsVIP          bool      `json:"is_vip" gorm:"default:false"`
	TotalQuestions int       `json:"total_questions" gorm:"column:question_count;default:0"`
	SortOrder      int       `json:"sort_order" gorm:"column:sort;default:0"`
	IsDeleted      bool      `json:"is_deleted" gorm:"default:false"`
	CreateTime     time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime     time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
}

func (QuestionBank) TableName() string {
	return "question_bank"
}

type Question struct {
	ID             int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title          string    `json:"title" gorm:"type:text;not null"`
	Content        string    `json:"content" gorm:"type:text"`
	Difficulty     int       `json:"difficulty" gorm:"type:smallint;default:2"`
	Category       string    `json:"category" gorm:"type:varchar(50)"`
	Tags           string    `json:"tags" gorm:"type:varchar(255)"`
	Answer         string    `json:"answer" gorm:"type:text"`
	Explanation    string    `json:"explanation" gorm:"type:text"`
	Heat           int       `json:"heat" gorm:"default:0"`
	BankID         int64     `json:"bank_id" gorm:"column:bank_id;default:0"`
	IsVIP          bool      `json:"is_vip" gorm:"default:false"`
	ViewCount      int       `json:"view_count" gorm:"default:0"`
	StarCount      int       `json:"star_count" gorm:"default:0"`
	LikeCount      int       `json:"like_count" gorm:"default:0"`
	AnsweredCount  int       `json:"answered_count" gorm:"default:0"`
	CorrectRate    int       `json:"correct_rate" gorm:"default:0"`
	IsDeleted      bool      `json:"is_deleted" gorm:"default:false"`
	CreateTime     time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime     time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
}

func (Question) TableName() string {
	return "question"
}

func (q *Question) GetDifficultyLabel() string {
	switch q.Difficulty {
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

type QuestionBankQuestion struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	BankID      int64     `json:"bank_id" gorm:"index:idx_bank_question;not null"`
	QuestionID  int64     `json:"question_id" gorm:"index:idx_bank_question;not null"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
}

func (QuestionBankQuestion) TableName() string {
	return "question_bank_question"
}

type QuestionTag struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"type:varchar(50);unique"`
	Category   string    `json:"category" gorm:"type:varchar(50)"`
	IsDeleted  bool      `json:"is_deleted" gorm:"default:false"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
}

func (QuestionTag) TableName() string {
	return "tag"
}

type QuestionTagRelation struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	QuestionID int64     `json:"question_id"`
	TagID      int64     `json:"tag_id"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
}

func (QuestionTagRelation) TableName() string {
	return "question_tag"
}
