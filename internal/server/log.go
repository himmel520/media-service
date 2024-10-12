package server

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func SetupLogger() *logrus.Logger {
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
	
	return log
}
