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
	adUC "github.com/himmel520/media-service/internal/usecase/ad"
	authUC "github.com/himmel520/media-service/internal/usecase/auth"
	colorUC "github.com/himmel520/media-service/internal/usecase/color"
	imgUC "github.com/himmel520/media-service/internal/usecase/img"
	tgUC "github.com/himmel520/media-service/internal/usecase/tg"

	"github.com/himmel520/media-service/internal/infrastructure/cache"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/postgres"
	adRepo "github.com/himmel520/media-service/internal/infrastructure/repository/postgres/ad"
	colorRepo "github.com/himmel520/media-service/internal/infrastructure/repository/postgres/color"
	imgRepo "github.com/himmel520/media-service/internal/infrastructure/repository/postgres/img"
	tgRepo "github.com/himmel520/media-service/internal/infrastructure/repository/postgres/tg"
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
	pool, err := postgres.NewPG(cfg.DB.DBConn)
	if err != nil {
		log.Fatalf("unable to connect to pool: %v", err)
	}
	defer pool.Close()
	dbtx := repository.NewDBTX(pool)

	// cache
	rdb, err := cache.NewRedis(cfg.Cache.Conn)
	if err != nil {
		log.Fatalf("unable to connect to cache: %v", err)
	}
	defer rdb.Close()

	cache := cache.Init(rdb, cfg.Cache.Exp)
	// repo := repository.New(db)
	// usecase := usecase.New(repo, cache, , log)

	// repo
	tgRepo := tgRepo.New()
	colorRepo := colorRepo.New()
	imgRepo := imgRepo.New()
	adRepo := adRepo.New()

	// uc
	authUC := authUC.New(cfg.Srv.JWT.PublicKey, log)
	tgUC := tgUC.New(dbtx, tgRepo, log)
	colorUC := colorUC.New(dbtx, colorRepo, log)
	imgUC := imgUC.New(dbtx, imgRepo, cache, log)
	adUC := adUC.New(dbtx, adRepo, cache, log)

	handler := ogen.NewHandler(ogen.HandlerParams{
		Auth:  authHandler.New(authUC, log),
		Error: errHandler.New(),
		Ad:    adHandler.New(adUC, log),
		Color: colorHandler.New(colorUC, log),
		Image: imgHandler.New(imgUC, log),
		Tg:    tgHandler.New(tgUC, log),
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
