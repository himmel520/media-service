package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/himmel520/media-service/config"
	"github.com/himmel520/media-service/internal/controller/ogen"
	"github.com/himmel520/media-service/internal/infrastructure/cache/redis"
	"github.com/himmel520/media-service/internal/infrastructure/repository/postgres"
	"github.com/himmel520/media-service/pkg/logger"
)

// @title API Documentation
// @version 1.0
// @description API для сервиса медиа и рекламы
// @host localhost:8081
// @BasePath /api/v1

func main() {
	// config
	logLevel := flag.String("loglevel", "info", "log level: debug, info, warn, error")
	flag.Parse()

	log := logger.SetupLogger(*logLevel)

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// db
	db, err := postgres.NewPG(cfg.DB.DBConn)
	if err != nil {
		log.Fatalf("unable to connect to pool: %v", err)
	}
	defer db.Close()

	// cache
	rdb, err := redis.NewRedis(cfg.Cache.Conn)
	if err != nil {
		log.Fatalf("unable to connect to cache: %v", err)
	}
	defer rdb.Close()

	// cache := cache.New(rdb, cfg.Cache.Exp)
	// repo := repository.New(db)
	// usecase := usecase.New(repo, cache, cfg.Srv.JWT.PublicKey, log)

	// handler := ogen.NewHandler(ogen.HandlerParams{
	// 	Auth:     auth.New(log)),
	// 	Error:    errHandler.New(),
	// 	Category: categoryHandler.New(categoryUC, log),
	// })

	app, err := ogen.NewServer(handler, cfg.Srv.Addr)
	if err != nil {
		log.Fatal(err)
	}

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
