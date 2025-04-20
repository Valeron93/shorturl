package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/Valeron93/shorturl/pkg/data"
)

type shorturlApi struct {
	db     data.UrlEntryTable
	logger *slog.Logger
}

func NewShorturlApi(db data.UrlEntryTable, logger *slog.Logger) shorturlApi {
	return shorturlApi{
		db:     db,
		logger: logger,
	}
}

func (s *shorturlApi) CreateNewUrl(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var request data.UrlEntry
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = s.db.CreateNewUrlEntry(request)
	if err != nil {
		http.Error(w, "failed to create url entry", http.StatusBadRequest)
		return
	}

	s.logger.Info("created url entry", "key", request.Key, "url", request.Url)

}

func (s *shorturlApi) Redirect(w http.ResponseWriter, r *http.Request) {

	url := r.PathValue("url")

	if url == "" {
		http.NotFound(w, r)
		return
	}

	urlEntry, err := s.db.GetUrlByKey(url)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, urlEntry.Url, http.StatusFound)
}
