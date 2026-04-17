package wire

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/d60-Lab/gin-template/internal/repository"
)

// RepositorySet 仓储层 Provider 集合
var RepositorySet = wire.NewSet(
	ProvideUserRepository,
	ProvidePracticeRouteRepository,
	ProvideQuestionRepository,
)

// ProvideUserRepository 提供用户仓储
func ProvideUserRepository(db *gorm.DB) repository.UserRepository {
	return repository.NewUserRepository(db)
}

// ProvidePracticeRouteRepository 提供刷题路线仓储
func ProvidePracticeRouteRepository(db *gorm.DB) repository.PracticeRouteRepository {
	return repository.NewPracticeRouteRepository(db)
}

// ProvideQuestionRepository 提供题目仓储
func ProvideQuestionRepository(db *gorm.DB) repository.QuestionRepository {
	return repository.NewQuestionRepository(db)
}
