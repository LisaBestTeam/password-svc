package config

import (
	"net"

	"github.com/spf13/viper"
)

func NetListener(getter *viper.Viper) net.Listener {
	addr := getter.GetString("net.addr")

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	return listen
}

func (c *config) Listener() net.Listener {
	return c.listener
}
