package config

import (
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/lisabestteam/password-svc/internal/horizon"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config interface {
	Log() *logrus.Logger
	Database() *sqlx.DB
	Horizon() horizon.HorizonClient
	Listener() net.Listener
}

type config struct {
	log      *logrus.Logger
	database *sqlx.DB
	client   horizon.HorizonClient
	listener net.Listener
}

func NewConfig(getter *viper.Viper) Config {
	return &config{
		database: NewDatabase(getter),
		log:      NewLog(getter),
		client:   NewHorizon(getter),
		listener: NetListener(getter),
	}
}

func MustGetter(configPath string) *viper.Viper {
	getter := viper.New()
	getter.SetConfigFile(configPath)
	if err := getter.ReadInConfig(); err != nil {
		panic(err)
	}

	return getter
}
