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
<<<<<<< HEAD
	DB1 *sql.DB // Supabase Database USER
	DB2 *sql.DB // Supabase Database MENU
	DB5 *sql.DB // Supabase Database ORDER
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
=======
	DB1        *sql.DB // Supabase Database USER
	DB2        *sql.DB // Supabase Database MENU
	AuthClient auth.Client
	Logger     *utils.Logger
}
>>>>>>> a45327588feb3cf73ade068eafcc07b4b38a9954
