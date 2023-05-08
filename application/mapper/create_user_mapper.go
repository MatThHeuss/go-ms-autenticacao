package mapper

import (
	"errors"
	"fmt"

	"github.com/MatThHeuss/go-user-microservice/internal/domain"
	"github.com/MatThHeuss/go-user-microservice/internal/domain/port"
)

func UserInputToUserEntity(input port.CreateUserInput) (*domain.User, error) {
	user, err := domain.NewUser(
		input.Name,
		input.Email,
		input.Birthday,
		input.Password,
		"user",
	)

	if err != nil {
		err := fmt.Sprintf("Error to create User. %s", err.Error())
		return nil, errors.New(err)
	}

	return user, nil
}

func UserEntityToCreateUserOutput(user *domain.User) *port.CreateUserOutput {
	return &port.CreateUserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Birthday:  user.Birthday,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.String(),
	}

}
