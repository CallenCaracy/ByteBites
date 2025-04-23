package graph

import "database/sql"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	DB1 *sql.DB // Supabase Database USER
	DB2 *sql.DB // Supabase Database MENU
	DB5 *sql.DB // Supabase Database ORDER
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
