package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/MatThHeuss/go-user-microservice/internal/domain/port"
	"go.uber.org/zap"
)

type CreateUserHandler struct {
	createUserUseCase port.CreateUserUseCase
	logger            *zap.Logger
}

func NewCreateUserHandler(createUserUseCase port.CreateUserUseCase, logger *zap.Logger) *CreateUserHandler {
	return &CreateUserHandler{
		createUserUseCase: createUserUseCase,
		logger:            logger,
	}
}

func (c *CreateUserHandler) Execute(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	payload := port.CreateUserInput{}
	c.logger.Info("create user Handler initiated")

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		c.logger.Error("Error to decode json to body", zap.Error(err))
		return err
	}

	output, err := c.createUserUseCase.Execute(ctx, payload)

	if err != nil {
		c.logger.Error("Error to execute use case in handler", zap.Error(err))
		return err
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(output)

}
