package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/MatThHeuss/go-user-microservice/internal/domain/port"
)

type CreateUserHandler struct {
	createUserUseCase port.CreateUserUseCase
}

func NewCreateUserHandler(createUserUseCase port.CreateUserUseCase) *CreateUserHandler {
	return &CreateUserHandler{
		createUserUseCase: createUserUseCase,
	}
}

func (c *CreateUserHandler) Execute(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	payload := port.CreateUserInput{}

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		return err
	}

	output, err := c.createUserUseCase.Execute(ctx, payload)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(output)

}
