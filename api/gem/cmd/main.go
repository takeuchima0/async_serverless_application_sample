package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/takeuchima0/async_serverless_application_sample/api/gem/internal/configuration"
	"github.com/takeuchima0/async_serverless_application_sample/api/gem/internal/controller"
)

func main() {
	ctx := context.Background()
	cnf, err := configuration.Load(ctx)
	if err != nil {
		slog.Error("Failed to read configuration", "error", err)
		panic(err)
	}

	server, err := controller.NewGemAPIServer(cnf)
	if err != nil {
		panic(err)
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(ctx, "failed to listen and serve", "error", err)
		}
	}()

	ctx, cancel := context.WithTimeout(ctx, 200*time.Microsecond)
	defer cancel()
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = server.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "shutdown server...", "error", err)
	}

	<-ctx.Done()
}
