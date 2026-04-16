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