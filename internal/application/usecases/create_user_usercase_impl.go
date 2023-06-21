package usecases

import (
	"context"
	"errors"
	"github.com/MatThHeuss/go-user-microservice/internal/application/mapper"
	"github.com/MatThHeuss/go-user-microservice/internal/domain/port"
	"github.com/MatThHeuss/go-user-microservice/internal/repository"
	"go.uber.org/zap"
)

type CreateUserUseCaseImpl struct {
	createUserRepository repository.UserRepository
	logger               *zap.Logger
}

func NewCreateUserUseCase(
	createUserRepository repository.UserRepository,
	logger *zap.Logger,
) port.CreateUserUseCase {
	return &CreateUserUseCaseImpl{
		createUserRepository: createUserRepository,
		logger:               logger,
	}
}

func (c *CreateUserUseCaseImpl) Execute(ctx context.Context, input port.CreateUserInput) (*port.CreateUserOutput, error) {
	c.logger.Info("Execute use case function")
	userEntity, err := mapper.UserInputToUserEntity(input)
	if err != nil {
		c.logger.Error("Error mapping user input to user entity", zap.Error(err))
		return nil, err
	}

	user, err := c.createUserRepository.FindByEmail(ctx, userEntity.Email)

	if err != nil {
		c.logger.Error("Error getting user", zap.Error(err))
		return nil, err
	}

	if user != nil {
		c.logger.Info("User already exists")
		return nil, errors.New("user already exists. Please login in your account")
	}

	err = c.createUserRepository.Create(ctx, userEntity)
	if err != nil {
		c.logger.Error("UseCase - error to create user repository", zap.Error(err))
		return nil, err
	}

	return mapper.UserEntityToCreateUserOutput(userEntity), nil
}
