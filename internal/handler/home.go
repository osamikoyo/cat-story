package handler

import (
	"net/http"

	"cat-story/internal/view/pages/home"
)

type serverHandler func(w http.ResponseWriter, r *http.Request) error

func Home(w http.ResponseWriter, r *http.Request) error {
	return home.HomePage().Render(r.Context(), w)
}
