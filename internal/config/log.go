package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLog(getter *viper.Viper) *logrus.Logger {
	level := getter.GetString("log.level")

	entry := logrus.New()

	switch level {
	case logrus.DebugLevel.String():
		entry.Level = logrus.DebugLevel
	//case logrus.InfoLevel.String():
	//	entry.Level = logrus.InfoLevel
	case logrus.WarnLevel.String():
		entry.Level = logrus.WarnLevel
	case logrus.ErrorLevel.String():
		entry.Level = logrus.ErrorLevel
	default:
		entry.Level = logrus.InfoLevel
	}

	return entry
}

func (c *config) Log() *logrus.Logger {
	return c.log
}
