package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/himmel520/media-service/config"
	"github.com/himmel520/media-service/internal/controller/ogen"
	adHandler "github.com/himmel520/media-service/internal/controller/ogen/ad"
	authHandler "github.com/himmel520/media-service/internal/controller/ogen/auth"
	colorHandler "github.com/himmel520/media-service/internal/controller/ogen/color"
	errHandler "github.com/himmel520/media-service/internal/controller/ogen/error"
	imgHandler "github.com/himmel520/media-service/internal/controller/ogen/img"
	tgHandler "github.com/himmel520/media-service/internal/controller/ogen/tg"

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

	handler := ogen.NewHandler(ogen.HandlerParams{
		Auth:  authHandler.New(log),
		Error: errHandler.New(),
		Ad:    adHandler.New(nil, log),
		Color: colorHandler.New(nil, log),
		Image: imgHandler.New(nil, log),
		Tg:    tgHandler.New(nil, log),
	})

	app, err := ogen.NewServer(handler, cfg.Srv.Addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Infof("the server is starting on %v", cfg.Srv.Addr)

		if err := app.Run(); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
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
