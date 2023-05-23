package main

import (
	"log"
	"net/http"

	"github.com/MatThHeuss/go-user-microservice/application/router"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Server Started",
		zap.String("Port", "8080"),
	)

	log.Fatal(http.ListenAndServe(":8080", router.SetupRoutes(logger)))
}
