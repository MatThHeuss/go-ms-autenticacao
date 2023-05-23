package router

import (
	"fmt"
	"net/http"

	"github.com/MatThHeuss/go-user-microservice/application/handler"
	"github.com/MatThHeuss/go-user-microservice/application/usecases"
	"github.com/MatThHeuss/go-user-microservice/internal/repository"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func SetupRoutes(logger *zap.Logger) http.Handler {

	r := chi.NewRouter()
	pgCliente, err := repository.NewPostgreSQLClient(logger)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	userRepository := repository.NewPostgreSQLUserRepository(pgCliente, logger)
	userUseCase := usecases.NewCreateUserUseCase(userRepository, logger)
	createUserHandler := handler.NewCreateUserHandler(userUseCase, logger)

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("/users called")
		err := createUserHandler.Execute(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return r

}
