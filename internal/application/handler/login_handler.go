package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/MatThHeuss/go-user-microservice/internal/domain/port"
	"go.uber.org/zap"
)

type LoginHandler struct {
	loginUseCase port.LoginUseCase
	logger       *zap.Logger
}

func NewLoginHandler(loginUseCase port.LoginUseCase, logger *zap.Logger) *LoginHandler {
	return &LoginHandler{
		loginUseCase: loginUseCase,
		logger:       logger,
	}
}

func (l *LoginHandler) Execute(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	payload := port.LoginInput{}
	l.logger.Info("create user Handler initiated")

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		l.logger.Error("Error to decode json to body", zap.Error(err))
		return err
	}

	output, err := l.loginUseCase.Execute(ctx, payload)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		l.logger.Error("Error to execute use case in handler", zap.Error(err))
		return json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(output)

}
