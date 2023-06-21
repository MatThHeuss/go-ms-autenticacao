package port

import "context"

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
}

type CreateUserOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type CreateUserUseCase interface {
	Execute(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error)
}
