package main

import (
	"log/slog"
	"net/http"

	"github.com/Valeron93/shorturl/pkg/api"
	"github.com/Valeron93/shorturl/pkg/middleware"
)

func main() {
	logger := slog.New(slog.Default().Handler())

	mux := http.NewServeMux()
	shorturlApi := api.NewShorturlApi()

	mux.HandleFunc("/hello", shorturlApi.Hello)
	handler := middleware.NewLoggingMiddleware(logger, mux)

	addr := ":8080"
	logger.Info("listening", "addr", addr)
	http.ListenAndServe(addr, handler)
}
