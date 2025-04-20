package middleware

import (
	"log/slog"
	"net/http"
)

func NewLoggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("request", "method", r.Method, "path", r.URL.Path, "addr", r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
