package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// "os"

	"auth-service/db"

	pb "auth-service/auth-service/proto"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

// HashPassword hashes a password
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

// Register a new user
func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	_, err = db.DB.Exec("INSERT INTO users (email, password_hash) VALUES ($1, $2)", req.Email, hashedPassword)
	if err != nil {
		return nil, err
	}

	fmt.Println("User registered:", req.Email)
	return &pb.AuthResponse{Token: "dummy-token"}, nil // Token will be generated later
}

func main() {
	db.InitDB()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen on port 50051:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &AuthServiceServer{})

	fmt.Println("Auth service is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve gRPC:", err)
	}
}
