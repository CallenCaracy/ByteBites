package main

import (
    "log"
    "net/http"
    "os"
    "strconv"

    "payment-service/db"
    "payment-service/graph"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/handler/extension"
    "github.com/99designs/gqlgen/graphql/handler/lru"
    "github.com/99designs/gqlgen/graphql/handler/transport"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"
const cacheSize = 100 // Define a reasonable cache size (e.g., 100 items)

func main() {
    // Initialize DB connection
    if err := db.InitDB(); err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    log.Println("✅ Connected to the database.")

    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    // Validate port
    if _, err := strconv.Atoi(port); err != nil {
        log.Fatalf("Invalid port number: %v. Please use a numeric port.", port)
    }

    // Initialize resolver (adjust if it requires dependencies like DB)
    resolver := graph.NewResolver(db.DB) // Update if NewResolver needs DB or other params

    // Set up GraphQL server
    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

    srv.AddTransport(transport.Options{})
    srv.AddTransport(transport.GET{})
    srv.AddTransport(transport.POST{})

    // Set query cache (for *ast.QueryDocument)
    queryCache := lru.New[*ast.QueryDocument](cacheSize)
    srv.SetQueryCache(queryCache)

    // Set automatic persisted query cache (likely maps string to *ast.QueryDocument or string)
    apqCache := lru.New[string](cacheSize) // Adjust type if needed (e.g., lru.New[*ast.QueryDocument])
    srv.Use(extension.Introspection{})
    srv.Use(extension.AutomaticPersistedQuery{
        Cache: apqCache,
    })

    // Set up handlers
    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    log.Printf("🚀 Server running at http://localhost:%s/", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}