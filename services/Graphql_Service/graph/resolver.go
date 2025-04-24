package graph

import (
	"database/sql"

	"github.com/CallenCaracy/ByteBites/services/User_Service/utils"
	"github.com/supabase-community/auth-go"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	DB1        *sql.DB // Supabase Database USER
	DB2        *sql.DB // Supabase Database MENU
	DB5        *sql.DB
	DB7        *sql.DB // Supabase Database KITCHEN
	AuthClient auth.Client
	Logger     *utils.Logger
}

// Mutation method that returns a mutation resolver.
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r} // Return the mutation resolver
}

// Query method that returns a query resolver.
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r} // Return the query resolver
}
