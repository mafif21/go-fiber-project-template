package app

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
)

func NewLogger(viper *viper.Viper) *logrus.Logger {
	log := logrus.New()
	logLevel, _ := strconv.Atoi(viper.GetString("LOG_LEVEL"))

	log.SetLevel(logrus.Level(logLevel))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
