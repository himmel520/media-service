package logger

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func SetupLogger(level string) *logrus.Logger {
	log := logrus.New()

	log.SetReportCaller(true)
	log.Formatter = &logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			s := strings.Split(f.Function, ".")
			fcname := s[len(s)-1]
			return fcname, fmt.Sprintf("%s:%d", f.File, f.Line)
		},
		PrettyPrint: true,
	}

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		log.Warn(err, "The level info is used")
		return log
	}

	log.Level = lvl

	return log
}
