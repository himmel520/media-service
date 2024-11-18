package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger(level string) {
	zlevel, err := zerolog.ParseLevel(level)
	if err != nil {
		zlevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(zlevel)
	zerolog.TimeFieldFormat = "02.01.2006 15:04:05"

	switch zlevel {
	case zerolog.DebugLevel:
		log.Logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "15:04:05"}).
			With().
			Timestamp().
			Caller().
			Logger()
	default:
		log.Logger = zerolog.New(os.Stdout).
			With().
			Timestamp().
			Caller().
			Logger()
	}
}