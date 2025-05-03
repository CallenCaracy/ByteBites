package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/CallenCaracy/ByteBites/services/User_Service/pb"
	"github.com/CallenCaracy/ByteBites/services/User_Service/utils"

	"os"

	"github.com/jackc/pgx/v5"
	"github.com/nedpals/supabase-go"
	"github.com/supabase-community/auth-go"
)

type UserServiceServer struct {
	pb.UnimplementedAuthServiceServer
	DB          *pgx.Conn
	AuthClient  auth.Client
	Logger      *utils.Logger
	SupabaseURL string
	APIKey      string
}

func (s *UserServiceServer) GetUserRole(ctx context.Context, req *pb.GetUserRoleRequest) (*pb.GetUserRoleResponse, error) {
	s.Logger.Info("Fetching user role from public.users for email: %s", req.Email)

	if req.Email == "" {
		return &pb.GetUserRoleResponse{Message: "user email cannot be empty"}, nil
	}

	query := `SELECT role FROM public.users WHERE email = $1`

	var role string
	row := s.DB.QueryRow(ctx, query, req.Email)
	err := row.Scan(&role)
	if err != nil {
		s.Logger.Error("Error fetching user role: %v", err)
		return &pb.GetUserRoleResponse{
			Message: "failed to fetch user role",
		}, err
	}

	return &pb.GetUserRoleResponse{
		Role:    role,
		Message: "Role retrieved successfully",
	}, nil
}

func (s *UserServiceServer) VerifyToken(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error) {
	s.Logger.Info("Verifying token...")

	supabaseURL := os.Getenv("SUPABASE_URL_FULL")
	supabaseKey := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		return nil, errors.New("supabase credentials are not set")
	}

	supabaseClient := supabase.CreateClient(supabaseURL, supabaseKey)

	userResponse, err := supabaseClient.Auth.User(ctx, req.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid or expired token: %v", err)
	}

	s.Logger.Info("Successfully verified token")

	return &pb.TokenResponse{
		Id:    userResponse.ID,
		Email: userResponse.Email,
	}, nil
}
