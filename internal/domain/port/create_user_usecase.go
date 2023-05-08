package port

import "context"

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
	Birthday string
}

type CreateUserOutput struct {
	ID        string
	Name      string
	Email     string
	Birthday  string
	Role      string
	CreatedAt string
}

type CreateUserUseCase interface {
	Execute(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error)
}
