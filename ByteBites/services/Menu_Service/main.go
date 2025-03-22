package main

import (
	"context"
	"net"
	"os"

	"Menu_Service/db"
	"Menu_Service/pb"
	menu_list "Menu_Service/server/menu_list"
	"Menu_Service/utils"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Initialize logger
	log := utils.NewLogger() //  Correct
	if log == nil {
		panic("Logger initialization failed")
	}

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: %v", err)
	}

	// Get the database URL from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// Connect to the database
	conn, err := db.ConnectDB(dbURL)
	if err != nil {
		log.Fatal("Database connection failed: %v", err)
	}
	defer func() {
		if err := conn.Close(context.Background()); err != nil {
			log.Fatal("Failed to close database connection: %v", err)
		}
	}()

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// ✅ Use the correct alias when creating the service
	menuService := &menu_list.MenuServiceServer{DB: conn, Logger: log}

	// Register the service with the gRPC server
	pb.RegisterMenuServiceServer(grpcServer, menuService)

	log.Info("Menu Service running on port 50053...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
