package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/d60-Lab/gin-template/internal/api/handler"
	"github.com/d60-Lab/gin-template/internal/api/middleware"
	"github.com/d60-Lab/gin-template/internal/dto"
	_ "github.com/d60-Lab/gin-template/openapi"
	"github.com/d60-Lab/gin-template/pkg/config"
)

func Setup(h *handler.Handler, aiHandler *handler.AIHandler, prHandler *handler.PracticeRouteHandler, questionHandler *handler.QuestionHandler, cfg *config.Config) *gin.Engine {
	r := gin.New()

	r.Use(middleware.CORS())
	r.Use(middleware.SecurityHeaders())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	if cfg.Sentry.Enabled {
		r.Use(middleware.Sentry())
	}

	if cfg.Tracing.Enabled {
		r.Use(middleware.Tracing(cfg.Tracing.ServiceName))
	}

	if cfg.Pprof.Enabled {
		middleware.Pprof(r)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/health", h.HealthCheck)

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", middleware.Validation(&dto.LoginRequest{}), h.Login)
			auth.GET("/me", middleware.Auth(cfg), h.GetCurrentUser)
			auth.PUT("/password", middleware.Auth(cfg), middleware.Validation(&dto.ChangePasswordRequest{}), h.ChangePassword)
			auth.GET("/stats", middleware.Auth(cfg), h.GetUserStats)
		}

		users := v1.Group("/users")
		{
			users.POST("", middleware.Validation(&dto.CreateUserRequest{}), h.CreateUser)
			users.GET("", middleware.Validation(&dto.ListUsersRequest{}), h.ListUsers)
			users.PUT("/me", middleware.Auth(cfg), h.UpdateCurrentUser)
			users.GET("/:id", middleware.Validation(&dto.GetUserRequest{}), h.GetUser)
			users.PUT("/:id", middleware.Auth(cfg), middleware.Validation(&dto.UpdateUserRequest{}), h.UpdateUser)
			users.DELETE("/:id", middleware.Auth(cfg), middleware.AdminOnly(), middleware.Validation(&dto.DeleteUserRequest{}), h.DeleteUser)
		}

		ai := v1.Group("/ai")
		{
			ai.POST("/verify", middleware.Validation(&dto.VerifyRequest{}), aiHandler.Verify)
		}

		practiceRoutes := v1.Group("/practice-routes")
		{
			practiceRoutes.GET("", prHandler.ListRoutes)
			practiceRoutes.GET("/:id", prHandler.GetRoute)
		}

		banks := v1.Group("/banks")
		{
			banks.GET("", middleware.Validation(&dto.ListBanksRequest{}), questionHandler.ListBanks)
			banks.GET("/:id", middleware.Validation(&dto.GetBankRequest{}), questionHandler.GetBank)
			banks.GET("/:id/questions", middleware.Validation(&dto.ListBankQuestionsRequest{}), questionHandler.ListBankQuestions)
		}

		questions := v1.Group("/questions")
		{
			questions.GET("", middleware.Validation(&dto.ListQuestionsRequest{}), questionHandler.ListQuestions)
			questions.GET("/hot", middleware.Validation(&dto.HotQuestionsRequest{}), questionHandler.ListHotQuestions)
			questions.GET("/:id", middleware.Validation(&dto.GetQuestionRequest{}), questionHandler.GetQuestion)
		}

		records := v1.Group("/records")
		{
			records.GET("/my", middleware.Auth(cfg), middleware.Validation(&dto.ListRecordsRequest{}), questionHandler.ListUserRecords)
			records.POST("", middleware.Auth(cfg), middleware.Validation(&dto.CreateRecordRequest{}), questionHandler.CreateRecord)
			records.PUT("/master", middleware.Auth(cfg), middleware.Validation(&dto.ToggleMasterRequest{}), questionHandler.ToggleMaster)
		}
	}

	return r
}
