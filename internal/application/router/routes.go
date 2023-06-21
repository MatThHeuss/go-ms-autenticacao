package router

import (
	"github.com/MatThHeuss/go-user-microservice/internal/application/handler"
	"github.com/MatThHeuss/go-user-microservice/internal/application/usecases"
	"net/http"

	"github.com/MatThHeuss/go-user-microservice/internal/repository"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func SetupRoutes(logger *zap.Logger) http.Handler {

	r := chi.NewRouter()
	pgClient, err := repository.NewPostgreSQLClient(logger)
	if err != nil {
		logger.Error("Error initializing PostgreSQL client", zap.Error(err))
		return nil
	}

	userRepository := repository.NewPostgreSQLUserRepository(pgClient, logger)
	userUseCase := usecases.NewCreateUserUseCase(userRepository, logger)
	createUserHandler := handler.NewCreateUserHandler(userUseCase, logger)

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("/users called")
		err := createUserHandler.Execute(w, r)
		if err != nil {
			logger.Error("Error executing createUserHandler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			logger.Info("createUserHandler executed successfully")
		}
	})

	return r

}