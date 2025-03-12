package main

import (
	"net"
	"os"

	"order-service/db"
	"order-service/pb"
	guestsession "order-service/server/guestsessions"
	order "order-service/server/orders"
	"order-service/utils"

	"context"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	log, err := utils.NewLogger()
	if err != nil {
		os.Exit(1)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: %v", err)
		os.Exit(1)
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
		os.Exit(1)
	}

	conn, err := db.ConnectDB(dbURL)
	if err != nil {
		log.Fatal("Database connection failed: %v", err)
		os.Exit(1)
	}
	defer func() {
		if err := conn.Close(context.Background()); err != nil {
			log.Fatal("Failed to close database connection: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	orderService := &order.OrderServiceServer{DB: conn}
	guestSessionService := &guestsession.GuestSessionServiceServer{DB: conn}
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	pb.RegisterGuestSessionServiceServer(grpcServer, guestSessionService)

	log.Info("Order Service running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
		os.Exit(1)
	}
}
