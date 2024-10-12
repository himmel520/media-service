package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/himmel520/uoffer/mediaAd/internal/config"
	httphandler "github.com/himmel520/uoffer/mediaAd/internal/handler/http"
	"github.com/himmel520/uoffer/mediaAd/internal/repository/postgres"
	"github.com/himmel520/uoffer/mediaAd/internal/repository/redis"
	"github.com/himmel520/uoffer/mediaAd/internal/server"
	"github.com/himmel520/uoffer/mediaAd/internal/service"
)

// @title API Documentation
// @version 1.0
// @description API для сервиса медиа и рекламы
// @host localhost:8081
// @BasePath /api/v1

func main() {
	log := server.SetupLogger()

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.NewPG(cfg.DB.DBConn)
	if err != nil {
		log.Fatalf("unable to connect to pool: %v", err)
	}

	cache, err := redis.New(&cfg.Cache)
	if err != nil {
		log.Fatalf("unable to connect to redis: %v", err)
	}

	repo := postgres.NewRepo(db)
	srv := service.New(repo, cache, log)
	handler := httphandler.New(srv, &cfg.Srv.JWT, log)

	// сервер

	app := server.New(handler.InitRoutes(), cfg.Srv.Addr)
	go func() {
		log.Infof("the server is starting on %v", cfg.Srv.Addr)

		if err := app.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occured while running http server: %s", err.Error())
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	if err := app.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err)
	}

	log.Info("the server is shut down")
}
