package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/idkOybek/newNewTerminal/pkg/auth"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
)

func AuthMiddleware(logger *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				logger.Error("Authorization header is missing")
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
				logger.Error("Invalid authorization header format")
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			token := bearerToken[1]
			claims, err := auth.ValidateToken(token)
			if err != nil {
				logger.Error("Invalid token", "error", err)
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AdminMiddleware(logger *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value("user").(*auth.Claims)
			if !ok || !claims.IsAdmin {
				logger.Error("User is not authorized as admin")
				http.Error(w, "Admin access required", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
