package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type PostgreSQLDbOperation interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

type PostgreSQLDbTransacion interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type PostgreSQLManager interface {
	Ping() error
}

type PostgreSQLClient interface {
	PostgreSQLDbOperation
	PostgreSQLDbTransacion
}

type Queries struct {
	Db     PostgreSQLDbOperation
	logger *zap.Logger
}

type PostgreSQLanitizedError struct {
	Code uint16
}

func (e PostgreSQLanitizedError) Error() string {
	return fmt.Sprintf("Database error code: %v. Check logs for more details.", e.Code)
}

func NewQueries(db PostgreSQLDbOperation, logger *zap.Logger) *Queries {
	return &Queries{Db: db, logger: logger}
}

func NewPostgreSQLClient(logger *zap.Logger) (PostgreSQLClient, error) {
	logger.Info("New postgresql client initiate")
	db, err := sql.Open("postgres", PostgreSQLConnectionString())
	if err != nil {
		logger.Error("error to open connection to postgresql", zap.Error(err))
		return nil, err
	}

	db.SetMaxOpenConns(1)
	db.SetMaxOpenConns(1)
	db.SetConnMaxLifetime(5)

	err = HealthCheck(db)
	if err != nil {
		db.Close()
		logger.Fatal("Failed to perform health check on PostgreSQL database", zap.Error(err))
		return nil, err
	}
	return db, nil
}

func PostgreSQLConnectionString() string {
	host := "localhost"
	port := 5432
	user := "root"
	password := "root"
	dbname := "ms"
	timeout := 5

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable connect_timeout=%v",
		host, port, user, password, dbname, timeout)

	return psqlInfo
}

func HealthCheck(db PostgreSQLManager) error {
	err := db.Ping()
	if err != nil {
		errHealthCheck := errors.New(fmt.Sprintf("Failed to perform health check operation on PostgreSQL database. %v", err.Error()))
		return errHealthCheck
	}

	return nil
}
