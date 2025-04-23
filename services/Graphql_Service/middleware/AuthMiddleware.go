package middleware

import (
	"net/http"

	"google.golang.org/grpc/metadata"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			// Create metadata with "authorization" key
			md := metadata.Pairs("authorization", authHeader)
			// Insert it into the request context
			ctx := metadata.NewIncomingContext(r.Context(), md)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}
