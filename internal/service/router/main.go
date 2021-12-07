package router

import (
	"context"
	"net"
	"net/http"
	"sync"

	"github.com/lisabestteam/password-svc/internal/config"
	"github.com/lisabestteam/password-svc/internal/database"
	"github.com/lisabestteam/password-svc/internal/database/postgres"
	"github.com/lisabestteam/password-svc/internal/service"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
)

func NewServer(cfg config.Config) service.Service {
	return &server{
		passwords: postgres.NewPassword(cfg.DB()),
		log:       cfg.Log(),
		listener:  cfg.Listener(),
		copus:     cfg.Copus(),
	}
}

type server struct {
	passwords database.Passwords
	log       *logan.Entry
	listener  net.Listener
	copus     types.Copus
}

func (r server) Run(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()

	log := r.log.WithField("service", "server")

	router := r.NewRouter()

	if err := r.copus.RegisterChi(router); err != nil {
		panic(err)
	}

	log.Info("Router init")

	if err := http.Serve(r.listener, router); err != nil {
		r.log.WithError(err).Error("failed to run server")
	}
}
