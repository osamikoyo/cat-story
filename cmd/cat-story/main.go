package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"cat-story/internal/app"
)

func main() {
	if err := realmain(); err != nil {
		slog.Error(err.Error())
	}
}

func realmain() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.App(ctx); err != nil{
		return err
	}

	return nil
}
