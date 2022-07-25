package log

import (
	"github.com/sirupsen/logrus"
	"github.com/thang-data/backend/config"
	"github.com/thang-data/backend/log/echologrus"
	"time"
)

func init() {
	// logrus
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
	cfg := config.GetConfig()
	level := ParseLevel(cfg)
	logrus.SetLevel(level)
	// echologrus
	echologrus.Logger().SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
	elevel := echologrus.ToEchoLevel(level)
	echologrus.Logger().SetLevel(elevel)
}

func ParseLevel(cfg config.ENV) logrus.Level {
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	return level
}
