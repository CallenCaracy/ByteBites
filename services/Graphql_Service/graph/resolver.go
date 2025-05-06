package graph

import (
	"Graphql_Service/graph/model"
	"database/sql"
	"sync"

	"github.com/CallenCaracy/ByteBites/services/User_Service/utils"
	"github.com/supabase-community/auth-go"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	DB1                        *sql.DB // Supabase Database USER
	DB2                        *sql.DB // Supabase Database MENU
	DB3                        *sql.DB // Supabase Database ORDER
	DB5                        *sql.DB
	DB7                        *sql.DB // Supabase Database KITCHEN
	AuthClient                 auth.Client
	Logger                     *utils.Logger
	MenuItemCreatedObservers   map[string]chan *model.MenuItem
	OrderQueueCreatedObservers map[string]chan *model.OrderQueue
	mu                         sync.Mutex
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }
