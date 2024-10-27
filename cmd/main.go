package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/himmel520/uoffer/mediaAd/config"
	httpctrl "github.com/himmel520/uoffer/mediaAd/internal/controller/http"

	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/cache/redis"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/mediaAd/internal/infrastructure/repository/postgres"
	"github.com/himmel520/uoffer/mediaAd/internal/server"
	"github.com/himmel520/uoffer/mediaAd/internal/usecase"
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
	defer db.Close()

	rdb, err := redis.New(cfg.Cache.Conn)
	if err != nil {
		log.Fatalf("unable to connect to cache: %v", err)
	}
	defer rdb.Close()

	cache := cache.New(rdb, cfg.Cache.Exp)
	repo := repository.New(db)
	usecase := usecase.New(repo, cache, cfg.Srv.JWT.PublicKey, log)

	handler := httpctrl.New(usecase, log)

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
