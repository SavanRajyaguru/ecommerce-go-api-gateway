package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ecommerce-go-api-gateway/api"
	"ecommerce-go-api-gateway/config"
	"ecommerce-go-api-gateway/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Init Logger
	logger.InitLogger(cfg.Logger.Level)
	defer logger.Log.Sync()
	logger.Log.Info("Starting API Gateway...")

	// 3. Setup Router
	r := api.SetupRouter(cfg)

	// 4. Start Server
	srv := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Listen: %s\n", zap.Error(err))
		}
	}()

	logger.Log.Info("Server started", zap.String("port", cfg.Server.Port))

	// 5. Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Server forced to shutdown:", zap.Error(err))
	}

	logger.Log.Info("Server exiting")
}
