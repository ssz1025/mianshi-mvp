package main

import (
	"fmt"
	"log"

	"github.com/d60-Lab/gin-template/internal/api/router"
	"github.com/d60-Lab/gin-template/internal/wire"
	"github.com/d60-Lab/gin-template/pkg/logger"
)

func main() {
	app, err := wire.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	cfg, err := wire.ProvideConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := logger.Init(cfg.Server.Mode); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	r := router.Setup(app.Handler, app.AIHandler, app.PracticeRouteHandler, app.QuestionHandler, cfg)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
