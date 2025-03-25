package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"Graphql_Service/pb"
	"Graphql_Service/utils"

	"strconv"
	"strings"

	"os"

	"github.com/jackc/pgx/v5"
	"github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
	"google.golang.org/grpc/metadata"
)

type UserServiceServer struct {
	pb.UnimplementedAuthServiceServer
	DB          *pgx.Conn
	AuthClient  auth.Client
	Logger      *utils.Logger
	SupabaseURL string
	APIKey      string
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

func (s *UserServiceServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	s.Logger.Info("Signing in %s", req.Email)

	signInData := types.SignupRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	authResponse, err := s.AuthClient.SignInWithEmailPassword(signInData.Email, signInData.Password)
	if err != nil {
		s.Logger.Error("Failed to sign in user: %v", err)
		return nil, fmt.Errorf("failed to sign in user: %v", err)
	}

	s.Logger.Info("User %s signed in successfully", authResponse.User.ID.String())

	return &pb.SignInResponse{
		AccessToken:  authResponse.AccessToken,
		RefreshToken: authResponse.RefreshToken,
		Error:        "",
	}, nil
}

// func (s *UserServiceServer) SignOut(ctx context.Context, req *pb.SignOutRequest) (*pb.SignOutResponse, error) {
// 	s.Logger.Info("Signing out user: %v", req.UserId)

// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return nil, fmt.Errorf("failed to retrieve metadata")
// 	}

// 	authHeader := md.Get("authorization")
// 	if len(authHeader) == 0 {
// 		return nil, fmt.Errorf("missing authorization token")
// 	}

// 	token := strings.TrimPrefix(authHeader[0], "Bearer ")

// 	s.AuthClient.WithToken(token)

// 	err := s.AuthClient.Logout()
// 	if err != nil {
// 		s.Logger.Error("Failed to sign out user: %v", err)
// 		return nil, fmt.Errorf("failed to sign out user: %v", err)
// 	}

// 	s.Logger.Info("User successfully signed out.")

// 	return &pb.SignOutResponse{
// 		Message: "User successfully signed out.",
// 		Error:   "",
// 	}, nil
// }

func (s *UserServiceServer) SignOut(ctx context.Context, req *pb.SignOutRequest) (*pb.SignOutResponse, error) {
	s.Logger.Info("Signing out user: %v", req.UserId)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		s.Logger.Error("Failed to retrieve metadata")
		return nil, fmt.Errorf("failed to retrieve metadata")
	}

	authHeader := md.Get("authorization")
	if len(authHeader) == 0 {
		s.Logger.Error("Missing authorization token")
		return nil, fmt.Errorf("missing authorization token")
	}
	accessToken := strings.TrimSpace(strings.TrimPrefix(authHeader[0], "Bearer "))

	refreshTokens := md.Get("refresh_token")
	refreshToken := ""
	if len(refreshTokens) > 0 {
		refreshToken = refreshTokens[0]
	}

	s.SupabaseURL = os.Getenv("SUPABASE_URL_FULL")
	serviceKey := os.Getenv("SERVICE_KEY")

	s.Logger.Info("Using Supabase URL: %s", s.SupabaseURL)
	s.Logger.Info("Using API Key: %s", serviceKey)

	if s.SupabaseURL == "" || serviceKey == "" {
		s.Logger.Error("Missing Supabase environment variables")
		return nil, fmt.Errorf("missing Supabase URL or API Key")
	}

	requestBody, err := json.Marshal(map[string]string{
		"refresh_token": refreshToken,
	})
	if err != nil {
		s.Logger.Error("Error marshaling JSON: %v", err)
		return nil, fmt.Errorf("failed to create logout request body: %v", err)
	}

	logoutURL := fmt.Sprintf("%s/auth/v1/logout", s.SupabaseURL)
	httpReq, err := http.NewRequest("POST", logoutURL, strings.NewReader(string(requestBody)))
	if err != nil {
		s.Logger.Error("Error creating logout request: %v", err)
		return nil, fmt.Errorf("failed to create logout request: %v", err)
	}

	httpReq.Header.Set("Authorization", "Bearer "+accessToken)
	httpReq.Header.Set("apikey", serviceKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		s.Logger.Error("Error sending logout request: %v", err)
		return nil, fmt.Errorf("failed to sign out user: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent { // 204 means successful logout
		s.Logger.Info("User successfully signed out.")
		return &pb.SignOutResponse{
			Message: "User successfully signed out.",
			Error:   "",
		}, nil
	}

	// Read response body (only if it's not 204)
	body, _ := io.ReadAll(resp.Body)
	s.Logger.Error("Failed to sign out user: Status %d, Response: %s", resp.StatusCode, string(body))
	return nil, fmt.Errorf("failed to sign out user: received status %d, response: %s", resp.StatusCode, string(body))
}

func (s *UserServiceServer) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	s.Logger.Info("Fetching user info from public.users for user_id: %s", req.UserId)

	if req.UserId == "" {
		return &pb.GetUserInfoResponse{Error: "user_id cannot be empty"}, nil
	}

	var user struct {
		UserID    string
		Email     string
		FirstName string
		LastName  string
		Role      string
		Address   string
		Phone     string
	}

	query := `SELECT id, email, first_name, last_name, role, address, phone FROM public.users WHERE id = $1`

	row := s.DB.QueryRow(ctx, query, req.UserId)
	err := row.Scan(&user.UserID, &user.Email, &user.FirstName, &user.LastName, &user.Role, &user.Address, &user.Phone)
	if err != nil {
		s.Logger.Error("Failed to fetch user info from public.users: %v", err)
		return &pb.GetUserInfoResponse{Error: "failed to fetch user info"}, nil
	}

	return &pb.GetUserInfoResponse{
		UserId:    user.UserID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		Address:   user.Address,
		Phone:     user.Phone,
	}, nil
}

func (s *UserServiceServer) UpdateUserInfo(ctx context.Context, req *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
	s.Logger.Info("Updating user info for user_id: %s", req.UserId)

	if req.UserId == "" {
		return nil, fmt.Errorf("user_id cannot be empty")
	}

	updates := []string{}
	args := []interface{}{req.UserId}

	if req.FirstName != nil {
		updates = append(updates, "first_name = $"+strconv.Itoa(len(args)+1))
		args = append(args, *req.FirstName)
	}
	if req.LastName != nil {
		updates = append(updates, "last_name = $"+strconv.Itoa(len(args)+1))
		args = append(args, *req.LastName)
	}
	if req.Role != nil {
		updates = append(updates, "role = $"+strconv.Itoa(len(args)+1))
		args = append(args, *req.Role)
	}
	if req.Address != nil {
		updates = append(updates, "address = $"+strconv.Itoa(len(args)+1))
		args = append(args, *req.Address)
	}
	if req.Phone != nil {
		updates = append(updates, "phone = $"+strconv.Itoa(len(args)+1))
		args = append(args, *req.Phone)
	}

	// If no fields to update, return an error
	if len(updates) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf("UPDATE public.users SET %s WHERE id = $1 RETURNING first_name, last_name, role, address, phone", strings.Join(updates, ", "))

	var updatedUser pb.UpdateUserInfoResponse
	row := s.DB.QueryRow(ctx, query, args...)
	err := row.Scan(&updatedUser.FirstName, &updatedUser.LastName, &updatedUser.Role, &updatedUser.Address, &updatedUser.Phone)
	if err != nil {
		s.Logger.Error("Failed to update user info for user_id %s: %v", req.UserId, err)
		return nil, fmt.Errorf("failed to update user info: %v", err)
	}

	s.Logger.Info("Successfully updated user info for user_id: %s", req.UserId)
	return &updatedUser, nil
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

func (s *UserServiceServer) ValidateToken(ctx context.Context) (*types.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	// Get the Authorization token
	authHeader := md.Get("authorization")
	if len(authHeader) == 0 {
		return nil, fmt.Errorf("missing authorization token")
	}

	token := strings.TrimPrefix(authHeader[0], "Bearer ")

	// Validate and parse the token
	s.AuthClient.WithToken(token)
	userResponse, err := s.AuthClient.GetUser()
	if err != nil {
		s.Logger.Error("Invalid or expired token: %v", err)
		return nil, fmt.Errorf("invalid or expired token")
	}

	return &userResponse.User, nil
}

func ExtractAuthToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("missing metadata")
	}

	authHeader := md.Get("authorization")
	if len(authHeader) == 0 {
		return "", fmt.Errorf("missing authorization token")
	}

	return authHeader[0], nil
}
