package main

import (
	"Graphql_Service/graph"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Get database URLs from .env
	db1URL := os.Getenv("SUPABASE_DB_USERS_URL")
	db2URL := os.Getenv("SUPABASE_DB_MENU_URL")
	db5URL := os.Getenv("SUPABASE_DB_ORDER_URL")

	// Connect to Supabase DB USERS
	db1, err := sql.Open("pgx", db1URL)
	if err != nil {
		log.Fatal("Failed to connect to Supabase DB1:", err)
	}
	defer db1.Close()

	// Connect to Supabase DB MENU
	db2, err := sql.Open("pgx", db2URL)
	if err != nil {
		log.Fatal("Failed to connect to Supabase DB2:", err)
	}
	defer db2.Close()

	// Connect to Supabase DB ORDER
	db5, err := sql.Open("pgx", db5URL)
	if err != nil {
		log.Fatal("Failed to connect to Supabase DB3:", err)
	}
	defer db5.Close()

	// Initialize GraphQL resolver with database connections
	resolver := &graph.Resolver{
		DB1: db1,
		DB2: db2,
		DB5: db5,
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
