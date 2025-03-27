package server

import (
	"Graphql_Service/db"
	"Graphql_Service/graph"
	"Graphql_Service/utils"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
)

// StartGraphQLServer runs the GraphQL server
func StartGraphQLServer(conn db.Querier) {
	logger, err := utils.NewLogger()
	if err != nil {
		logger.Fatal("Failed to create logger: %v", err)
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: conn}}))

	// Add WebSocket transport to support subscriptions (or WebSocket requests)
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			// Allow all origins for simplicity; adjust for production
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		KeepAlivePingInterval: 10 * time.Second,
	})

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	logger.Info("🚀 GraphQL Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start GraphQL server: ", err)
	}
}
