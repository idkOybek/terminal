package middleware

import (
	"net/http"
	"time"

	"github.com/idkOybek/newNewTerminal/pkg/logger"
)

func LoggerMiddleware(log *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Call the next handler
			next.ServeHTTP(w, r)

			// Log the request
			log.Infow("HTTP request",
				"method", r.Method,
				"path", r.URL.Path,
				"duration", time.Since(start),
				"remote_addr", r.RemoteAddr,
			)
		})
	}
}
