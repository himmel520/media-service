package log

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func SetupLogger(level string) {
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
	}

	log.Level = lvl
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}