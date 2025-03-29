package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.70

import (
	"Graphql_Service/graph/model"
	"context"
	"database/sql"
	"fmt"

	"github.com/CallenCaracy/ByteBites/services/User_Service/pb"
	"google.golang.org/grpc"
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
	panic(fmt.Errorf("not implemented: SignIn - signIn"))
}

// SignInOnlyEmployee is the resolver for the signInOnlyEmployee field.
func (r *mutationResolver) SignInOnlyEmployee(ctx context.Context, input model.SignInEmployeeInput) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented: SignInOnlyEmployee - signInOnlyEmployee"))
}

// SignOut is the resolver for the signOut field.
func (r *mutationResolver) SignOut(ctx context.Context) (bool, error) {
	panic(fmt.Errorf("not implemented: SignOut - signOut"))
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
