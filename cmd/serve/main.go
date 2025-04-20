package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/Valeron93/shorturl/pkg/api"
	"github.com/Valeron93/shorturl/pkg/data"
	"github.com/Valeron93/shorturl/pkg/middleware"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	logger := slog.New(slog.Default().Handler())

	mux := http.NewServeMux()

	dbString := "./db.sql"

	db, err := sql.Open("sqlite", "./db.sqlite")
	if err != nil {
		logger.Error("failed to connect to DB", "addr", dbString, "err", err)
		os.Exit(1)
	}
	logger.Info("connected to db", "addr", dbString)

	urlEntryTable, err := data.NewUrlEntryTable(db)
	if err != nil {
		logger.Error("failed to create UrlEntryTable", "err", err)
	}

	shorturlApi := api.NewShorturlApi(urlEntryTable, logger)

	mux.HandleFunc("POST /url/", shorturlApi.CreateNewUrl)
	mux.HandleFunc("/r/{url}", shorturlApi.Redirect)
	handler := middleware.NewLoggingMiddleware(logger, mux)

	addr := ":8080"
	logger.Info("listening", "addr", addr)
	http.ListenAndServe(addr, handler)
}
