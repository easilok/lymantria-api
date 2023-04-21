package config

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func SetupLogger() {
	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(&logrus.TextFormatter{
		// DisableColors: true,
		FullTimestamp: true,
	})

	// log.SetReportCaller(true)
}
