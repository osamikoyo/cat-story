package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Dependensis struct{
	AssetsFS http.FileSystem
}

func RegisterHandler(r *chi.Mux, deps Dependensis) {
	r.Get("/", handlerErrRoute(Home))
	r.Get("/get-stories", handlerErrRoute(GetAll))

	r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(deps.AssetsFS)))
}

func handlerErrRoute(h serverHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil{
			slog.Error(err.Error())
		}
	}
}