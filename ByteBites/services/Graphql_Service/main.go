package main

import (
	"net"
	"os"

	"Graphql_Service/db"
	"Graphql_Service/pb"
	user "Graphql_Service/server/users"
	"Graphql_Service/utils"

	"context"

	"github.com/joho/godotenv"
	"github.com/supabase-community/auth-go"
	"google.golang.org/grpc"
)

func main() {
	log, err := utils.NewLogger()
	if err != nil {
		log.Fatal("Failed to create logger: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: %v", err)
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANON_KEY")
	if supabaseURL == "" || supabaseKey == "" {
		log.Fatal("Supabase URL or Anon Key environment variable not set")
	}

	client := auth.New(supabaseURL, supabaseKey)
	if client == nil {
		log.Fatal("Failed to create auth client")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	conn, err := db.ConnectDB(dbURL)
	if err != nil {
		log.Fatal("Database connection failed: %v", err)
	}
	defer func() {
		if err := conn.Close(context.Background()); err != nil {
			log.Fatal("Failed to close database connection: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userService := &user.UserServiceServer{DB: conn, AuthClient: client, Logger: log}
	pb.RegisterAuthServiceServer(grpcServer, userService)

	log.Info("Order Service running on port 50050...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
