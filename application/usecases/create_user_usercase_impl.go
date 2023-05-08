package usecases

import (
	"context"

	"github.com/MatThHeuss/go-user-microservice/application/mapper"
	"github.com/MatThHeuss/go-user-microservice/internal/domain/port"
	"github.com/MatThHeuss/go-user-microservice/internal/repository"
)

type CreateUserUseCaseImpl struct {
	createUserRepository repository.UserRepository
}

func NewCreateUserUseCase(
	createUserRepository repository.UserRepository,
) port.CreateUserUseCase {
	return &CreateUserUseCaseImpl{
		createUserRepository: createUserRepository,
	}
}

func (c *CreateUserUseCaseImpl) Execute(ctx context.Context, input port.CreateUserInput) (*port.CreateUserOutput, error) {

	userEntity, err := mapper.UserInputToUserEntity(input)
	if err != nil {
		return nil, err
	}

	err = c.createUserRepository.Create(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	return mapper.UserEntityToCreateUserOutput(userEntity), nil
}
