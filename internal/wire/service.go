package wire

import (
	"github.com/google/wire"

	"github.com/d60-Lab/gin-template/internal/repository"
	"github.com/d60-Lab/gin-template/internal/service"
	"github.com/d60-Lab/gin-template/pkg/config"
)

// ServiceSet 服务层 Provider 集合
var ServiceSet = wire.NewSet(
	ProvideUserService,
	ProvideAIService,
	ProvidePracticeRouteService,
	ProvideQuestionService,
)

// ProvideUserService 提供用户服务
func ProvideUserService(
	userRepo repository.UserRepository,
	cfg *config.Config,
) service.UserService {
	return service.NewUserService(userRepo, cfg)
}

// ProvideAIService 提供 AI 服务
func ProvideAIService(
	modelClients map[string]service.AIModelClient,
) service.AIService {
	return service.NewAIService(modelClients)
}

// ProvidePracticeRouteService 提供刷题路线服务
func ProvidePracticeRouteService(
	practiceRouteRepo repository.PracticeRouteRepository,
) service.PracticeRouteService {
	return service.NewPracticeRouteService(practiceRouteRepo)
}

// ProvideQuestionService 提供题目服务
func ProvideQuestionService(
	questionRepo repository.QuestionRepository,
) service.QuestionService {
	return service.NewQuestionService(questionRepo)
}