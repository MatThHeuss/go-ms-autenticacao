package main

import (
	"github.com/MatThHeuss/go-user-microservice/internal/application/router"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	http.ListenAndServe(":8080", router.SetupRoutes(logger))
	logger.Info("Server Started",
		zap.String("Port", "8080"),
	)
}
