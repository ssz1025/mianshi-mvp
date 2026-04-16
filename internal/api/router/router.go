package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/d60-Lab/gin-template/internal/api/handler"
	"github.com/d60-Lab/gin-template/internal/api/middleware"
	"github.com/d60-Lab/gin-template/internal/dto"
	_ "github.com/d60-Lab/gin-template/openapi" // swagger docs
	"github.com/d60-Lab/gin-template/pkg/config"
)

func Setup(h *handler.Handler, cfg *config.Config) *gin.Engine {
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
		}

		users := v1.Group("/users")
		{
			users.POST("", middleware.Validation(&dto.CreateUserRequest{}), h.CreateUser)
			users.GET("", middleware.Validation(&dto.ListUsersRequest{}), h.ListUsers)
			users.GET("/:id", middleware.Validation(&dto.GetUserRequest{}), h.GetUser)
			users.PUT("/:id", middleware.Auth(cfg), middleware.Validation(&dto.UpdateUserRequest{}), h.UpdateUser)
			users.DELETE("/:id", middleware.Auth(cfg), middleware.AdminOnly(), middleware.Validation(&dto.DeleteUserRequest{}), h.DeleteUser)
		}
	}

	return r
}
