package middleware

import (
	"context"
	"fmt"
	"net/http"
)

func HelloMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "data-rahasia", "rahasia")
		fmt.Println("Hello from middleware")
		next.ServeHTTP(w, r.WithContext(ctx))
		fmt.Println("Hello again from middleware")
	})
}
