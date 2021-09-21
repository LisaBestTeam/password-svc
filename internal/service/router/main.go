package router

import (
	"context"
	"github.com/lisabestteam/password-svc/internal/database/postgres"
	"net"
	"net/http"
	"sync"

	"github.com/lisabestteam/password-svc/internal/config"
	"github.com/lisabestteam/password-svc/internal/database"
	"github.com/lisabestteam/password-svc/internal/service"
	"github.com/sirupsen/logrus"
)

type Router interface {
	service.Service
}

func NewRouter(cfg config.Config) Router {
	return router{
		passwords: postgres.NewPassword(cfg.Database()),
		log:       cfg.Log(),
		listener:  cfg.Listener(),
	}
}

type router struct {
	passwords database.Passwords
	log       *logrus.Logger
	listener  net.Listener
}

func (r router) Run(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()

	log := r.log.WithField("service", "router")

	router := r.NewRouter()

	log.Info("Router init")

	if err := http.Serve(r.listener, router); err != nil {
		r.log.WithError(err).Error("failed to run serve")
	}
}
