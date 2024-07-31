// internal/middleware/logger_middleware.go

package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Call the next handler
			next.ServeHTTP(w, r)

			// Log the request
			logger.Info("Request",
				zap.String("method", r.Method),
				zap.String("uri", r.RequestURI),
				zap.String("addr", r.RemoteAddr),
				zap.Duration("duration", time.Since(start)),
			)
		})
	}
}
