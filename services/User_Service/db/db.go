package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Querier interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Role      string
	Address   string
	Phone     string
	IsActive  string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func ConnectDB(dbURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
