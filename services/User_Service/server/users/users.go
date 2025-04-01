package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/CallenCaracy/ByteBites/services/User_Service/pb"
	"github.com/CallenCaracy/ByteBites/services/User_Service/utils"

	"strings"

	"os"

	"github.com/golang-jwt/jwt/v4"
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

func (s *UserServiceServer) SignInOnlyEmployee(ctx context.Context, req *pb.SignInOnlyEmployeeRequest) (*pb.SignInOnlyEmployeeResponse, error) {
	s.Logger.Info("Attempting to sign in %s", req.Email)

	roleResp, err := s.GetUserRole(ctx, &pb.GetUserRoleRequest{Email: req.Email})
	if err != nil {
		s.Logger.Error("Error retrieving role for %s: %v", req.Email, err)
		return nil, fmt.Errorf("failed to retrieve user role: %v", err)
	}

	if roleResp.Role != "employee" {
		s.Logger.Error("Email %s has role %s; not allowed to sign in as employee", req.Email, roleResp.Role)
		return nil, fmt.Errorf("user does not have permission to sign in as an employee")
	}

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

	return &pb.SignInOnlyEmployeeResponse{
		AccessToken:  authResponse.AccessToken,
		RefreshToken: authResponse.RefreshToken,
		Error:        "",
	}, nil
}

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

func (s *UserServiceServer) GetUserRole(ctx context.Context, req *pb.GetUserRoleRequest) (*pb.GetUserRoleResponse, error) {
	s.Logger.Info("Fetching user info from public.users for email: %s", req.Email)

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
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, errors.New("JWT secret key is not set")
	}

	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	id, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid user ID in token")
	}

	email, _ := claims["email"].(string)

	return &pb.TokenResponse{
		Id:    id,
		Email: email,
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
