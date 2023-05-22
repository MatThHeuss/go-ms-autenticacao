package router

import (
	"fmt"
	"net/http"

	"github.com/MatThHeuss/go-user-microservice/application/handler"
	"github.com/MatThHeuss/go-user-microservice/application/usecases"
	"github.com/MatThHeuss/go-user-microservice/internal/repository"
	"github.com/go-chi/chi"
)

func SetupRoutes() http.Handler {

	r := chi.NewRouter()
	pgCliente, err := repository.NewPostgreSQLClient()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	userRepository := repository.NewPostgreSQLUserRepository(pgCliente)
	userUseCase := usecases.NewCreateUserUseCase(userRepository)
	createUserHandler := handler.NewCreateUserHandler(userUseCase)

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		err := createUserHandler.Execute(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return r

}
