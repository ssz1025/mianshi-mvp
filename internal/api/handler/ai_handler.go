package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/d60-Lab/gin-template/internal/api/middleware"
	"github.com/d60-Lab/gin-template/internal/dto"
	"github.com/d60-Lab/gin-template/internal/service"
	"github.com/d60-Lab/gin-template/pkg/response"
)

type AIHandler struct {
	aiService service.AIService
}

func NewAIHandler(aiService service.AIService) *AIHandler {
	return &AIHandler{
		aiService: aiService,
	}
}

func (h *AIHandler) Verify(c *gin.Context) {
	req := middleware.MustGetRequest[dto.VerifyRequest](c)

	summary, responses, err := h.aiService.GenerateAnswers(c.Request.Context(), req.Question, req.Models)
	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Success(c, gin.H{
		"summary":   summary,
		"responses": responses,
	})
}