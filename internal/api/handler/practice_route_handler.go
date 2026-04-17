package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/d60-Lab/gin-template/internal/service"
	"github.com/d60-Lab/gin-template/pkg/response"
)

type PracticeRouteHandler struct {
	practiceRouteService service.PracticeRouteService
}

func NewPracticeRouteHandler(practiceRouteService service.PracticeRouteService) *PracticeRouteHandler {
	return &PracticeRouteHandler{
		practiceRouteService: practiceRouteService,
	}
}

func (h *PracticeRouteHandler) ListRoutes(c *gin.Context) {
	category := c.Query("category")

	result, err := h.practiceRouteService.ListRoutes(c.Request.Context(), category)
	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Success(c, result)
}

func (h *PracticeRouteHandler) GetRoute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(c, "invalid route id")
		return
	}

	result, err := h.practiceRouteService.GetRoute(c.Request.Context(), id)
	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Success(c, result)
}
