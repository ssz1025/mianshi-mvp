package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/d60-Lab/gin-template/internal/api/middleware"
	"github.com/d60-Lab/gin-template/internal/dto"
	"github.com/d60-Lab/gin-template/internal/service"
	"github.com/d60-Lab/gin-template/pkg/response"
)

type QuestionHandler struct {
	questionService service.QuestionService
}

func NewQuestionHandler(questionService service.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		questionService: questionService,
	}
}

func (h *QuestionHandler) ListBanks(c *gin.Context) {
	req := middleware.MustGetRequest[dto.ListBanksRequest](c)
	req.SetDefaults()

	result, err := h.questionService.ListBanks(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Success(c, result)
}

func (h *QuestionHandler) GetBank(c *gin.Context) {
	req := middleware.MustGetRequest[dto.GetBankRequest](c)

	bank, err := h.questionService.GetBankByID(c.Request.Context(), req.ID)
	if err != nil {
		if err == service.ErrBankNotFound {
			response.NotFound(c, "bank not found")
			return
		}
		response.InternalError(c, err)
		return
	}

	response.Success(c, bank)
}

func (h *QuestionHandler) ListBankQuestions(c *gin.Context) {
	req := middleware.MustGetRequest[dto.ListBankQuestionsRequest](c)
	req.SetDefaults()

	result, err := h.questionService.ListBankQuestions(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Success(c, result)
}

func (h *QuestionHandler) GetQuestion(c *gin.Context) {
	req := middleware.MustGetRequest[dto.GetQuestionRequest](c)

	question, err := h.questionService.GetQuestionByID(c.Request.Context(), req.ID)
	if err != nil {
		if err == service.ErrQuestionNotFound {
			response.NotFound(c, "question not found")
			return
		}
		response.InternalError(c, err)
		return
	}

	response.Success(c, question)
}

func (h *QuestionHandler) ListHotQuestions(c *gin.Context) {
	req := middleware.MustGetRequest[dto.HotQuestionsRequest](c)
	req.SetDefaults()

	result, err := h.questionService.ListHotQuestions(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Success(c, result)
}

func (h *QuestionHandler) ListQuestions(c *gin.Context) {
	req := middleware.MustGetRequest[dto.ListQuestionsRequest](c)
	req.SetDefaults()

	result, err := h.questionService.ListQuestions(c.Request.Context(), req)
	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Success(c, result)
}
