package repository

import (
	"context"

	"github.com/MatThHeuss/go-user-microservice/internal/domain"
)

type PostgreSQLUserRepository struct {
	*Queries
	Db PostgreSQLClient
}

func NewPostgreSQLUserRepository(postgreSQLClient PostgreSQLClient) UserRepository {
	postgreSQLUserRepository := &PostgreSQLUserRepository{
		Db:      postgreSQLClient,
		Queries: NewQueries(postgreSQLClient),
	}

	return postgreSQLUserRepository
}

func (u *PostgreSQLUserRepository) Create(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5)
	defer cancel()

	err := u.execTx(ctx, func(q *Queries) error {
		var err error
		err = q.insert(ctx, user)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

func (q *Queries) insert(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5)
	defer cancel()

	query := "INSERT INTO users (id, name, birthday, email, password, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := q.Db.ExecContext(ctx, query,
		user.ID,
		user.Name,
		user.Birthday,
		user.Email,
		user.Password,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil

}

func (u *PostgreSQLUserRepository) execTx(ctx context.Context, fn func(*Queries) error, tags ...string) error {
	tx, txErr := u.Db.BeginTx(ctx, nil)
	if txErr != nil {
		return txErr
	}

	q := NewQueries(tx)
	err := fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}

		return err
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		return commitErr
	}

	return nil
}
