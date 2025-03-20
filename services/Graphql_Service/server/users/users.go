package server

import (
	"context"
	"fmt"

	"Graphql_Service/pb"
	"Graphql_Service/utils"

	"github.com/jackc/pgx/v5"
	"github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
)

type UserServiceServer struct {
	pb.UnimplementedAuthServiceServer
	DB         *pgx.Conn
	AuthClient auth.Client
	Logger     *utils.Logger
}

func (s *UserServiceServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	s.Logger.Info("Received signup request for email: %s", req.Email)

	signUpData := types.SignupRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := s.AuthClient.Signup(signUpData)
	if err != nil {
		s.Logger.Error("Failed to sign up user: %v", err)
		return nil, fmt.Errorf("failed to sign up user: %v", err)
	}

	s.Logger.Info("Successfully signed up user: %s", user.ID.String())

	var phone, address *string
	if req.Phone != nil && *req.Phone != "" {
		phone = req.Phone
	}
	if req.Address != nil && *req.Address != "" {
		address = req.Address
	}

	_, err = s.DB.Exec(ctx, `
		INSERT INTO users (id, email, first_name, last_name, role, address, phone)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		user.ID, req.Email, req.FirstName, req.LastName, req.Role, address, phone,
	)
	if err != nil {
		s.Logger.Error("Failed to insert user into database: %v", err)
		return nil, fmt.Errorf("failed to insert user into database: %v", err)
	}

	s.Logger.Info("User %s successfully inserted into the database", user.ID.String())

	return &pb.SignUpResponse{
		UserId:    user.ID.String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
	}, nil
}

func (s *UserServiceServer) DeactivateUser(ctx context.Context, req *pb.DeactivateUserRequest) (*pb.DeactivateUserResponse, error) {
	s.Logger.Info("Received request to deactivate user: %s", req.UserId)

	var isActive string
	err := s.DB.QueryRow(ctx, `SELECT is_active FROM users WHERE id = $1`, req.UserId).Scan(&isActive)
	if err != nil {
		s.Logger.Error("Failed to check user status: %v", err)
		return nil, fmt.Errorf("failed to check user status: %v", err)
	}

	if isActive == "inactive" {
		return &pb.DeactivateUserResponse{
			Message: "User is already deactivated",
		}, nil
	}

	_, err = s.DB.Exec(ctx, `UPDATE users SET is_active = 'inactive', updated_at = NOW() WHERE id = $1`, req.UserId)
	if err != nil {
		s.Logger.Error("Failed to deactivate user: %v", err)
		return nil, fmt.Errorf("failed to deactivate user: %v", err)
	}

	s.Logger.Info("User %s successfully deactivated", req.UserId)

	return &pb.DeactivateUserResponse{
		Message: "User deactivated successfully",
	}, nil
}

func (s *UserServiceServer) ReactivateUser(ctx context.Context, req *pb.ReactivateUserRequest) (*pb.ReactivateUserResponse, error) {
	s.Logger.Info("Received request to reactivate user: %s", req.UserId)

	var isActive string
	err := s.DB.QueryRow(ctx, `SELECT is_active FROM users WHERE id = $1`, req.UserId).Scan(&isActive)
	if err != nil {
		s.Logger.Error("Failed to check user status: %v", err)
		return nil, fmt.Errorf("failed to check user status: %v", err)
	}

	if isActive == "active" {
		return &pb.ReactivateUserResponse{
			Message: "User is already active",
		}, nil
	}

	_, err = s.DB.Exec(ctx, `UPDATE users SET is_active = 'active', updated_at = NOW() WHERE id = $1`, req.UserId)
	if err != nil {
		s.Logger.Error("Failed to reactivate user: %v", err)
		return nil, fmt.Errorf("failed to reactivate user: %v", err)
	}

	s.Logger.Info("User %s successfully reactivated", req.UserId)

	return &pb.ReactivateUserResponse{
		Message: "User reactivated successfully",
	}, nil
}
