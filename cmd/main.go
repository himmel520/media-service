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

	log "github.com/youroffer/logger"
	
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/postgres"
	adRepo "github.com/himmel520/media-service/internal/infrastructure/repository/postgres/ad"
	colorRepo "github.com/himmel520/media-service/internal/infrastructure/repository/postgres/color"
	imgRepo "github.com/himmel520/media-service/internal/infrastructure/repository/postgres/img"
	tgRepo "github.com/himmel520/media-service/internal/infrastructure/repository/postgres/tg"
)

func init() {
	logLevel := flag.String("loglevel", "info", "log level: debug, info, warn, error")
	flag.Parse()

	log.SetupLogger(*logLevel)
}

func main() {
	// config
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// db
	pool, err := postgres.NewPG(cfg.DB.DBConn)
	if err != nil {
		log.FatalMsg(err, "unable to connect to pool")
	}
	defer pool.Close()
	dbtx := repository.NewDBTX(pool)

	// cache
	rdb, err := cache.NewRedis(cfg.Cache.Conn)
	if err != nil {
		log.FatalMsg(err, "unable to connect to cache")
	}
	defer rdb.Close()
	cache := cache.NewCache(rdb, cfg.Cache.Exp)

	// repo
	tgRepo := tgRepo.New()
	colorRepo := colorRepo.New()
	imgRepo := imgRepo.New()
	adRepo := adRepo.New()

	// uc
	authUC := authUC.New(cfg.Srv.JWT.PublicKey)
	tgUC := tgUC.New(dbtx, tgRepo, cache)
	colorUC := colorUC.New(dbtx, colorRepo, cache)
	imgUC := imgUC.New(dbtx, imgRepo, cache)
	adUC := adUC.New(dbtx, adRepo, cache)

	handler := ogen.NewHandler(ogen.HandlerParams{
		Auth:  authHandler.New(authUC),
		Error: errHandler.New(),
		Ad:    adHandler.New(adUC),
		Color: colorHandler.New(colorUC),
		Image: imgHandler.New(imgUC),
		Tg:    tgHandler.New(tgUC),
	})

	app, err := ogen.NewServer(handler, cfg.Srv.Addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Infof("the server is starting on %v", cfg.Srv.Addr)

		if err := app.Run(); err != nil {
			log.FatalMsg(err, "error occured while running http server")
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	if err := app.Shutdown(context.Background()); err != nil {
		log.ErrMsg(err, "error occured on server shutting down")
	}

	log.Info("the server is shut down")
}