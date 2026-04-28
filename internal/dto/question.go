package dto

type ListBanksRequest struct {
	Category string `form:"category" binding:"omitempty"`
	Search   string `form:"search" binding:"omitempty"`
	Page     int    `form:"page" binding:"omitempty,gte=1"`
	PageSize int    `form:"page_size" binding:"omitempty,gte=1,lte=100"`
}

func (r *ListBanksRequest) SetDefaults() {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.PageSize < 1 || r.PageSize > 100 {
		r.PageSize = 20
	}
}

type GetBankRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type ListQuestionsRequest struct {
	BankID     int64  `form:"bank_id" binding:"omitempty"`
	Category   string `form:"category" binding:"omitempty"`
	Difficulty string `form:"difficulty" binding:"omitempty"`
	Search     string `form:"search" binding:"omitempty"`
	SortBy     string `form:"sort_by" binding:"omitempty"`
	Tag        string `form:"tag" binding:"omitempty"`
	Page       int    `form:"page" binding:"omitempty,gte=1"`
	PageSize   int    `form:"page_size" binding:"omitempty,gte=1,lte=100"`
}

func (r *ListQuestionsRequest) SetDefaults() {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.PageSize < 1 || r.PageSize > 100 {
		r.PageSize = 30
	}
}

type GetQuestionRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type ListBankQuestionsRequest struct {
	ID         int64  `uri:"id" binding:"required"`
	Difficulty string `form:"difficulty" binding:"omitempty"`
	Search     string `form:"search" binding:"omitempty"`
	Page       int    `form:"page" binding:"omitempty,gte=1"`
	PageSize   int    `form:"page_size" binding:"omitempty,gte=1,lte=100"`
}

func (r *ListBankQuestionsRequest) SetDefaults() {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.PageSize < 1 || r.PageSize > 100 {
		r.PageSize = 50
	}
}

type HotQuestionsRequest struct {
	Category   string `form:"category" binding:"omitempty"`
	Difficulty string `form:"difficulty" binding:"omitempty"`
	Tag        string `form:"tag" binding:"omitempty"`
	SortBy     string `form:"sort_by" binding:"omitempty"`
	Page       int    `form:"page" binding:"omitempty,gte=1"`
	PageSize   int    `form:"page_size" binding:"omitempty,gte=1,lte=100"`
}

func (r *HotQuestionsRequest) SetDefaults() {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.PageSize < 1 || r.PageSize > 100 {
		r.PageSize = 30
	}
}

type QuestionBankResponse struct {
	ID             int64  `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Icon           string `json:"icon"`
	Category       string `json:"category"`
	IsVIP          bool   `json:"is_vip"`
	TotalQuestions int    `json:"total_questions"`
	UpdateTime     string `json:"update_time"`
}

type QuestionListItem struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Category    string   `json:"category"`
	Difficulty  string   `json:"difficulty"`
	Tags        []string `json:"tags"`
	IsVIP       bool     `json:"is_vip"`
	ViewCount   int      `json:"view_count"`
	StarCount   int      `json:"star_count"`
	LikeCount   int      `json:"like_count"`
}

type QuestionDetail struct {
	ID               int64                `json:"id"`
	Title            string               `json:"title"`
	Content          string               `json:"content"`
	Category         string               `json:"category"`
	Difficulty       string               `json:"difficulty"`
	Tags             []string             `json:"tags"`
	IsVIP            bool                 `json:"is_vip"`
	ViewCount        int                  `json:"view_count"`
	StarCount        int                  `json:"star_count"`
	LikeCount        int                  `json:"like_count"`
	Answer           string               `json:"answer"`
	Explanation      string               `json:"explanation"`
	RelatedQuestions []RelatedQuestion    `json:"related_questions"`
}

type RelatedQuestion struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type ListRecordsRequest struct {
	Filter   string `form:"filter" binding:"omitempty,oneof=all master not-master"`
	Page     int    `form:"page" binding:"omitempty,gte=1"`
	PageSize int    `form:"page_size" binding:"omitempty,gte=1,lte=100"`
}

func (r *ListRecordsRequest) SetDefaults() {
	if r.Filter == "" {
		r.Filter = "all"
	}
	if r.Page < 1 {
		r.Page = 1
	}
	if r.PageSize < 1 || r.PageSize > 100 {
		r.PageSize = 10
	}
}

type QuestionRecordItem struct {
	ID                int64    `json:"id"`
	QuestionID        int64    `json:"question_id"`
	QuestionTitle     string   `json:"question_title"`
	QuestionDifficulty string  `json:"question_difficulty"`
	QuestionCategory  string   `json:"question_category"`
	QuestionTags      []string `json:"question_tags"`
	IsMaster          bool     `json:"is_master"`
	LastViewTime      string   `json:"last_view_time"`
	CreateTime        string   `json:"create_time"`
}

type PaginatedResponse struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

type ToggleFavoriteRequest struct {
	QuestionID int64 `json:"question_id" binding:"required"`
}

type FavoriteItem struct {
	ID             int64    `json:"id"`
	QuestionID     int64    `json:"question_id"`
	QuestionTitle  string   `json:"question_title"`
	QuestionCategory string `json:"question_category"`
	QuestionDifficulty string `json:"question_difficulty"`
	QuestionTags   []string `json:"question_tags"`
	CreateTime     string   `json:"create_time"`
}

type CreateRecordRequest struct {
	QuestionID int64 `json:"question_id" binding:"required"`
}

type ToggleMasterRequest struct {
	QuestionID int64 `json:"question_id" binding:"required"`
	IsMaster   bool  `json:"is_master"`
}
