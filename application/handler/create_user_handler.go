package handler

import (
	"fmt"
	"net/http"

	"github.com/MatThHeuss/go-user-microservice/internal/domain/port"
)

type CreateUserHandler struct {
	createUserUseCase port.CreateUserUseCase
}

func (c *CreateUserHandler) Execute(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Hello world")
	return nil
}
