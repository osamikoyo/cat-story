package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"cat-story/internal/config"
	"cat-story/internal/handler"
)

func App(ctx context.Context) error {
	cfg := config.New()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	s := http.Server{
		Addr:    cfg.Addr,
		Handler: r,
	}

	handler.RegisterHandler(r, handler.Dependensis{AssetsFS: http.Dir(cfg.Assets)})

	go func() {
		<-ctx.Done()
		slog.Info("shotdown server")
		s.Shutdown(ctx)
	}()

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed{
		return err
	}

	return nil
}
