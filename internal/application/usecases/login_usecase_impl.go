package usecases

import (
	"context"
	"errors"
	"github.com/MatThHeuss/go-user-microservice/internal/auth"
	"github.com/MatThHeuss/go-user-microservice/internal/domain/port"
	"github.com/MatThHeuss/go-user-microservice/internal/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type LoginImplementation struct {
	userRepository repository.UserRepository
	logger         *zap.Logger
}

func NewLoginUseCase(userRepository repository.UserRepository, logger *zap.Logger) port.LoginUseCase {
	return &LoginImplementation{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (l LoginImplementation) Execute(ctx context.Context, input port.LoginInput) (*port.LoginOutput, error) {
	l.logger.Info("Login usecase started")

	user, err := l.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		l.logger.Error("Error getting user", zap.Error(err))
		return nil, errors.New("invalid credentials")
	}

	if user == nil {
		l.logger.Info("User not found")
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	accessToken := auth.CreateToken(user)

	return &port.LoginOutput{
		AccessToken: accessToken,
	}, nil
}
