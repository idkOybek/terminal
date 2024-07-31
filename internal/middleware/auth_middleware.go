// internal/middleware/auth_middleware.go

package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/idkOybek/newNewTerminal/pkg/auth"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
	"go.uber.org/zap"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		token := bearerToken[1]

		claims, err := auth.ValidateToken(token)
		if err != nil {
			logger.Error("Invalid token", zap.Error(err))
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
