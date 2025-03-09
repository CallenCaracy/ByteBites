package utils

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor is a gRPC unary interceptor that checks for a valid JWT token.
// Public methods (Register, Login, Recover) are exempted.
func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// List of public methods that do not require authentication.
	publicMethods := map[string]bool{
		"/auth.AuthService/Register": true,
		"/auth.AuthService/Login":    true,
		"/auth.AuthService/Recover":  true,
	}

	if publicMethods[info.FullMethod] {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata not provided")
	}

	authHeaders := md.Get("authorization")
	if len(authHeaders) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization token not provided")
	}

	tokenStr := authHeaders[0]
	// Remove "Bearer " prefix if present.
	if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
		tokenStr = tokenStr[7:]
	}

	email, err := ValidateJWT(tokenStr)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	// Optionally, add the email to the context for downstream handlers.
	type contextKey string
	const userEmailKey contextKey = "userEmail"
	newCtx := context.WithValue(ctx, userEmailKey, email)
	return handler(newCtx, req)
}

// newGRPCServer creates a gRPC server with the authentication interceptor.
func NewGRPCServer() *grpc.Server {
	return grpc.NewServer(
		grpc.UnaryInterceptor(AuthInterceptor),
	)
}
