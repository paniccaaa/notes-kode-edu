package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/paniccaaa/notes-kode-edu/internal/config"
	"github.com/paniccaaa/notes-kode-edu/internal/http/router"
	"github.com/paniccaaa/notes-kode-edu/internal/lib/logger"
	noteservice "github.com/paniccaaa/notes-kode-edu/internal/services/note-service"
	"github.com/paniccaaa/notes-kode-edu/internal/storage/postgres"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	storage, err := postgres.NewPostgres(cfg.DbURI)
	if err != nil {
		log.Error("failed to init storage", slog.String("err", err.Error()))
		os.Exit(1)
	}

	defer storage.Db.Close()

	noteService := noteservice.NewNoteService(storage)
	router := router.InitRouter(log, noteService)

	log.Info("starting notes-kode-edu", slog.String("addr", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed to start server", slog.String("err", err.Error()))
		}
	}()

	log.Info("server started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", slog.String("err", err.Error()))
		return
	}

	log.Info("server stopped")

}
