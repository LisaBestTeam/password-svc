package config

import (
	"errors"

	"github.com/lisabestteam/password-svc/internal/horizon"
	"github.com/spf13/viper"
)

var errEndpointHorizonEmpty = errors.New("horizon endpoint can`t be is empty")

func NewHorizon(getter *viper.Viper) horizon.HorizonClient {
	endpoint := getter.GetString("horizon.endpoint")
	if endpoint == "" {
		panic(errEndpointHorizonEmpty)
	}

	return horizon.NewHorizonClient(endpoint)
}

func (c config) Horizon() horizon.HorizonClient {
	return c.client
}
