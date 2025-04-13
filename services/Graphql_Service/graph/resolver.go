package graph

import (
	"Graphql_Service/db"
	"Graphql_Service/pb"
	"context"
	"fmt"
	"time"
	// "github.com/jackc/pgx/v5"
)

type Resolver struct {
	DB db.Querier
}

func (r *Resolver) Query_getUser(ctx context.Context, id string) (*pb.User, error) {
	var user db.User
	err := r.DB.QueryRow(ctx, `
    SELECT 
        id, email, first_name, last_name, role, address, phone, is_active, created_at, updated_at
    FROM users
    WHERE id = $1
	`, id).Scan(
		&user.ID, &user.Email, &user.FirstName, &user.LastName,
		&user.Role, &user.Address, &user.Phone, &user.IsActive,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	updatedAtStr := ""
	if user.UpdatedAt != nil {
		updatedAtStr = user.UpdatedAt.Format(time.RFC3339)
	}

	return &pb.User{
		Id:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		Address:   user.Address,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedAtStr,
	}, nil
}
