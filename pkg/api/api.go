package api

import (
	"fmt"
	"net/http"
)

type shorturlApi struct {
}

func NewShorturlApi() shorturlApi {
	return shorturlApi{}
}

func (s *shorturlApi) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
