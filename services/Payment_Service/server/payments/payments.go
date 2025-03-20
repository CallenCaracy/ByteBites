package server

import (
	"github.com/jackc/pgx/v5"
)

type OrderServiceServer struct {
	DB *pgx.Conn
}

// Your logic here
