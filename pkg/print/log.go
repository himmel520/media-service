package print

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	Logger.SetReportCaller(true)
	Logger.Formatter = &logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			s := strings.Split(f.Function, ".")
			fcname := s[len(s)-1]
			return fcname, fmt.Sprintf("%s:%d", f.File, f.Line)
		},
		PrettyPrint: true,
	}
}

func SetupLogger(level string) *logrus.Logger {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		Logger.Infoln(err, "The level info is used")
		return Logger
	}

	Logger.Level = lvl

	return Logger
}
