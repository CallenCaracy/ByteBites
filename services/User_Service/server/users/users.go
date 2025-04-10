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

// For authenticated Users
// func (s *UserServiceServer) ChangeUserPassword(ctx context.Context, req *pb.ChangeUserPasswordRequest) (*pb.ChangeUserPasswordResponse, error) {
// 	s.Logger.Info("Changing password for user: %s", req.UserId)

// 	// Validate the new password
// 	if req.NewPassword == "" {
// 		return nil, fmt.Errorf("new password cannot be empty")
// 	}

// 	// Retrieve the Authorization token from metadata
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		s.Logger.Error("No metadata found in context")
// 		return nil, fmt.Errorf("failed to retrieve metadata from context")
// 	}

// 	// Debug metadata content
// 	s.Logger.Info("Metadata received: %+v", md)

// 	authHeader := md.Get("authorization")
// 	if len(authHeader) == 0 {
// 		s.Logger.Error("Authorization header is missing")
// 		return nil, fmt.Errorf("missing authorization token")
// 	}

// 	// Extract Bearer token
// 	tokenParts := strings.Split(authHeader[0], " ")
// 	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
// 		s.Logger.Error("Invalid authorization format: %v", authHeader[0])
// 		return nil, fmt.Errorf("invalid authorization header format")
// 	}
// 	accessToken := tokenParts[1]

// 	// Log the token for debugging (Avoid in production)
// 	s.Logger.Info("Extracted Token: %s", accessToken)

// 	// Prepare the update data
// 	updateData := types.UpdateUserRequest{
// 		Password: &req.NewPassword,
// 	}

// 	// Call UpdateUser with the update data
// 	user, err := s.AuthClient.UpdateUser(updateData) // Only one parameter is passed
// 	if err != nil {
// 		s.Logger.Error("Failed to change password for user %s: %v", req.UserId, err)
// 		return nil, fmt.Errorf("failed to change password: %v", err)
// 	}

// 	s.Logger.Info("Password successfully changed for user: %s", user.ID.String())

// 	return &pb.ChangeUserPasswordResponse{
// 		Message: "Password changed successfully",
// 	}, nil
// }

// We'll come back to this
// func (s *UserServiceServer) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
// 	s.Logger.Info("Requesting password reset for email: %s", req.Email)

// 	recoverRequest := types.RecoverRequest{
// 		Email: req.Email,
// 	}
// 	err := s.AuthClient.Recover(recoverRequest)
// 	if err != nil {
// 		s.Logger.Error("Failed to send password reset email to %s: %v", req.Email, err)
// 		return nil, fmt.Errorf("failed to send password reset email: %v", err)
// 	}

// 	s.Logger.Info("Password reset email sent to: %s", req.Email)

// 	return &pb.ForgotPasswordResponse{
// 		Message: "Password reset email sent successfully",
// 	}, nil
// }

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
