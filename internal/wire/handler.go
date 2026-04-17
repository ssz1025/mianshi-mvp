package wire

import (
	"github.com/google/wire"

	"github.com/d60-Lab/gin-template/internal/api/handler"
	"github.com/d60-Lab/gin-template/internal/model"
	"github.com/d60-Lab/gin-template/internal/service"
)

// HandlerSet Handler 层 Provider 集合
var HandlerSet = wire.NewSet(
	ProvideHandler,
	ProvideAIHandler,
	ProvidePracticeRouteHandler,
	ProvideQuestionHandler,
)

// ProvideHandler 提供主 Handler
func ProvideHandler(userService service.UserService) *handler.Handler {
	return handler.NewHandler(userService)
}

// ProvideAIHandler 提供 AI Handler
func ProvideAIHandler(aiService service.AIService) *handler.AIHandler {
	return handler.NewAIHandler(aiService)
}

// ProvidePracticeRouteHandler 提供刷题路线 Handler
func ProvidePracticeRouteHandler(practiceRouteService service.PracticeRouteService) *handler.PracticeRouteHandler {
	return handler.NewPracticeRouteHandler(practiceRouteService)
}

// ProvideQuestionHandler 提供题目 Handler
func ProvideQuestionHandler(questionService service.QuestionService) *handler.QuestionHandler {
	return handler.NewQuestionHandler(questionService)
}

// ProvideAIModelClients 提供 AI 模型客户端集合
func ProvideAIModelClients() map[string]service.AIModelClient {
	return map[string]service.AIModelClient{
		"qwen":   model.NewQwenClient(),
		"gpt":    model.NewGPTClient(),
		"claude": model.NewClaudeClient(),
	}
}