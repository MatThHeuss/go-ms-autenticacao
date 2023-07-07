package repository

import (
	"context"
	"github.com/MatThHeuss/go-user-microservice/internal/domain"
	"go.uber.org/zap"
)

type PostgreSQLUserRepository struct {
	*Queries
	Db PostgreSQLClient
}

func NewPostgreSQLUserRepository(postgreSQLClient PostgreSQLClient, logger *zap.Logger) UserRepository {
	postgreSQLUserRepository := &PostgreSQLUserRepository{
		Db:      postgreSQLClient,
		Queries: NewQueries(postgreSQLClient, logger),
	}

	return postgreSQLUserRepository
}

func (u *PostgreSQLUserRepository) Create(ctx context.Context, user *domain.User) error {
	//ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	//defer cancel()

	err := u.execTx(ctx, func(q *Queries) error {
		var err error
		err = q.insert(ctx, user)
		if err != nil {
			u.logger.Error("Error in create query", zap.Error(err))
			return err
		}

		return nil
	})
	u.logger.Error("Error in create query", zap.Error(err))
	return err
}

func (u *PostgreSQLUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	//ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	//defer cancel()
	query := "SELECT id, name, birthday, email, password, role, created_at, updated_at FROM users WHERE email = $1"

	row := u.Db.QueryRowContext(
		ctx, query, email,
	)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.Birthday, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		u.logger.Error("Error scaning user", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

func (q *Queries) insert(ctx context.Context, user *domain.User) error {
	q.logger.Info("Insert query started")

	query := "INSERT INTO users (id, name, birthday, email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"

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
		q.logger.Error("Error in insert query", zap.Error(err))
		return err
	}

	return nil

}

func (u *PostgreSQLUserRepository) execTx(ctx context.Context, fn func(*Queries) error, tags ...string) error {
	tx, txErr := u.Db.BeginTx(ctx, nil)
	if txErr != nil {
		u.logger.Info("Error to initiate transaction", zap.Error(txErr))
		return txErr
	}

	q := NewQueries(tx, u.logger)
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
