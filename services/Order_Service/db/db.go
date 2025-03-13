package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(dbURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
