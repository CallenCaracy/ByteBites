package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.70

import (
	"Graphql_Service/graph/model"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/CallenCaracy/ByteBites/services/User_Service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpInput) (*model.User, error) {
	conn, err := grpc.Dial("localhost:50050", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	req := &pb.SignUpRequest{
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Role:      input.Role,
	}

	if input.Address != nil {
		req.Address = input.Address
	}
	if input.Phone != nil {
		req.Phone = input.Phone
	}

	res, err := client.SignUp(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to sign up user: %v", err)
	}

	return &model.User{
		ID:        res.UserId,
		Email:     input.Email,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Role:      res.Role,
		Address:   input.Address,
		Phone:     input.Phone,
	}, nil
}

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.AuthResponse, error) {
	conn, err := grpc.Dial("localhost:50050", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	req := &pb.SignInRequest{
		Email:    input.Email,
		Password: input.Password,
	}

	res, err := client.SignIn(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to login user: %v", err)
	}

	return &model.AuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		Error:        &res.Error,
	}, nil
}

// SignInOnlyEmployee is the resolver for the signInOnlyEmployee field.
func (r *mutationResolver) SignInOnlyEmployee(ctx context.Context, input model.SignInEmployeeInput) (*model.AuthResponse, error) {
	conn, err := grpc.Dial("localhost:50050", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	req := &pb.SignInOnlyEmployeeRequest{
		Email:    input.Email,
		Password: input.Password,
	}

	res, err := client.SignInOnlyEmployee(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to login user: %v", err)
	}

	return &model.AuthResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		Error:        &res.Error,
	}, nil
}

// SignOut is the resolver for the signOut field.
func (r *mutationResolver) SignOut(ctx context.Context) (bool, error) {
	// Extract HTTP headers from the request
	// requestCtx := graphql.GetOperationContext(ctx)
	// if requestCtx == nil {
	//     return false, fmt.Errorf("missing HTTP request context")
	// }

	// httpRequest, ok := requestCtx.RawRequest.(*http.Request)
	// if !ok {
	//     return false, fmt.Errorf("failed to extract HTTP request")
	// }

	// authToken := httpRequest.Header.Get("Authorization")
	// refreshToken := httpRequest.Header.Get("refresh_token")

	// if authToken == "" || refreshToken == "" {
	//     return false, fmt.Errorf("missing Authorization or refresh_token headers")
	// }

	// // Debugging logs
	// fmt.Println("Extracted Authorization:", authToken)
	// fmt.Println("Extracted Refresh Token:", refreshToken)

	// // Attach headers as metadata to gRPC request
	// md := metadata.New(map[string]string{
	//     "authorization": authToken,
	//     "refresh_token": refreshToken,
	// })
	// ctx = metadata.NewOutgoingContext(ctx, md)

	// // Establish gRPC connection
	// conn, err := grpc.Dial("localhost:50050", grpc.WithInsecure())
	// if err != nil {
	//     return false, fmt.Errorf("failed to connect to gRPC server: %v", err)
	// }
	// defer conn.Close()

	// client := pb.NewAuthServiceClient(conn)

	// // Make the gRPC call
	// res, err := client.SignOut(ctx, &pb.SignOutRequest{})
	// if err != nil {
	//     return false, fmt.Errorf("sign out failed: %v", err)
	// }

	// if res.Error != "" {
	//     return false, fmt.Errorf("sign out error: %s", res.Error)
	// }

	// return true, nil
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// GetUserByID - Fetch a user by ID
func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User

	query := `SELECT id, email, first_name, last_name, role, address, phone, is_active, created_at, updated_at FROM public.users WHERE id = $1`
	err := r.Resolver.DB1.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.FirstName, &user.LastName,
		&user.Role, &user.Address, &user.Phone, &user.IsActive,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if user not found
		}
		return nil, err
	}
	return &user, nil
}

// GetAuthenticatedUser is the resolver for the getAuthenticatedUser field.
func (r *queryResolver) GetAuthenticatedUser(ctx context.Context) (*model.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata in context")
	}
	log.Printf("Received metadata: %+v\n", md)

	tokenList := md.Get("authorization")
	if len(tokenList) == 0 {
		return nil, fmt.Errorf("missing token in metadata")
	}

	token := strings.TrimPrefix(tokenList[0], "Bearer ")

	conn, err := grpc.Dial("localhost:50050", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	resp, err := client.VerifyToken(ctx, &pb.TokenRequest{Token: token})
	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %v", err)
	}

	user, err := r.GetUserByID(ctx, resp.Id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &model.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		Address:   user.Address,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
