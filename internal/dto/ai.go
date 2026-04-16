package dto

type VerifyRequest struct {
	QuestionID string   `json:"questionId" binding:"required"`
	Question   string   `json:"question" binding:"required"`
	Models     []string `json:"models" binding:"omitempty"`
}

type ModelResponse struct {
	Content string `json:"content"`
	Model   string `json:"model"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type VerifyResponse struct {
	Summary   string                   `json:"summary"`
	Responses map[string]ModelResponse `json:"responses"`
}