package main

import (
	"fmt"
	"log"

	"github.com/d60-Lab/gin-template/internal/api/router"
	"github.com/d60-Lab/gin-template/internal/wire"
	"github.com/d60-Lab/gin-template/pkg/config"
	"github.com/d60-Lab/gin-template/pkg/logger"
)

func main() {
	if err := logger.Init("debug"); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	app, err := wire.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := router.Setup(app.Handler, app.AIHandler, app.PracticeRouteHandler, app.QuestionHandler, cfg)
	log.Printf("Server starting on port %d", cfg.Server.Port)
	if err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
