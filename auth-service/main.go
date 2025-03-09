package main

import (
	"fmt"
	"log"
	"net"

	pb "auth-service/auth-service/proto"
	"auth-service/internal/db"
	"auth-service/internal/server"
	"auth-service/internal/utils"
)

func main() {
	db.InitDB()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen on port 50051:", err)
	}

	// Create a new gRPC server with the authentication interceptor.
	grpcServer := utils.NewGRPCServer()

	pb.RegisterAuthServiceServer(grpcServer, server.NewAuthServiceServer())

	fmt.Println("Auth service is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve gRPC:", err)
	}
}
