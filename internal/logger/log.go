package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var Log *logrus.Logger

func Init() {

	Log = logrus.New()
	logLevel := ""
	envstr := os.Getenv("LOG_LEVEL")
	if envstr != "" {
		logLevel = strings.ToUpper(envstr)
		Log.Infof("Setting log level to: %s\n", envstr)
	}

	switch logLevel {
	case "TRACE":
		Log.SetLevel(logrus.TraceLevel)
	case "DEBUG":
		Log.SetLevel(logrus.DebugLevel)
	case "INFO":
		Log.SetLevel(logrus.InfoLevel)
	case "WARN":
		Log.SetLevel(logrus.WarnLevel)
	case "ERROR":
		Log.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		Log.SetLevel(logrus.FatalLevel)
	case "PANIC":
		Log.SetLevel(logrus.PanicLevel)
	default:
		Log.SetLevel(logrus.ErrorLevel)
	}

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

}
