package server

import (
	"github.com/jackc/pgx/v5"
)

type PaymentServiceServer struct {
	DB *pgx.Conn
}

// Your logic here
