package server

import (
	"context"
	"fmt"

	pb "auth-service/auth-service/proto"
	"auth-service/internal/db"
	"auth-service/internal/utils"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServiceServer implements the AuthService gRPC service.
type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

// NewAuthServiceServer returns a new instance of AuthServiceServer.
func NewAuthServiceServer() *AuthServiceServer {
	return &AuthServiceServer{}
}

// HashPassword hashes a given password using bcrypt.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

// Register creates a new user.
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
	// For now, return a dummy token (later you might automatically log the user in)
	return &pb.AuthResponse{Token: "dummy-token"}, nil
}

// Login validates the user's credentials and returns a JWT token.
func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	var storedHash string
	err := db.DB.QueryRow("SELECT password_hash FROM users WHERE email = $1", req.Email).Scan(&storedHash)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	token, err := utils.GenerateJWT(req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not generate token")
	}

	return &pb.AuthResponse{Token: token}, nil
}

// Recover initiates a password recovery process and returns a reset token.
func (s *AuthServiceServer) Recover(ctx context.Context, req *pb.RecoverRequest) (*pb.RecoverResponse, error) {
	var email string
	err := db.DB.QueryRow("SELECT email FROM users WHERE email = $1", req.Email).Scan(&email)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	resetToken, err := utils.GenerateResetToken(req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not generate reset token")
	}

	return &pb.RecoverResponse{Token: resetToken}, nil
}
