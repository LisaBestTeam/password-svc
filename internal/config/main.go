package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Listenerer
	pgdb.Databaser
	types.Copuser
	comfig.Logger
}

type config struct {
	comfig.Listenerer
	pgdb.Databaser
	types.Copuser
	comfig.Logger

	getter kv.Getter
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		Listenerer: comfig.NewListenerer(getter),
		Databaser:  pgdb.NewDatabaser(getter),
		Copuser:    copus.NewCopuser(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		getter:     getter,
	}
}
