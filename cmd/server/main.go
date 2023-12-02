package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"initial-http-server/internal/config"
	"initial-http-server/internal/handler"
	"initial-http-server/internal/logger"
)

func main() {
	// init configuration
	appConfig := config.InitConfig()

	// init logger
	logger.InitLogger(appConfig.Logging)
	defer logger.Logger.Sync()

	// init gin router
	router := gin.Default()

	// define routes
	router.GET("/hello", handler.HelloHandler)

	port := appConfig.Server.Port
	addr := fmt.Sprintf(":%d", port)

	logger.Logger.Info("Server starting...", zap.String("address", addr))
	if err := router.Run(addr); err != nil {
		logger.Logger.Fatal("Server failed to start", zap.Error(err))
	}
}
