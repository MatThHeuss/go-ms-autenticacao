package repository

import (
	"context"

	"github.com/MatThHeuss/go-user-microservice/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}
