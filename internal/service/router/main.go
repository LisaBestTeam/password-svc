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
	"github.com/sirupsen/logrus"
)

func NewServer(cfg config.Config) service.Service {
	return &server{
		passwords: postgres.NewPassword(cfg.Database()),
		log:       cfg.Log(),
		listener:  cfg.Listener(),
	}
}

type server struct {
	passwords database.Passwords
	log       *logrus.Logger
	listener  net.Listener
}

func (r server) Run(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()

	log := r.log.WithField("service", "server")

	router := r.NewRouter()

	log.Info("Router init")

	if err := http.Serve(r.listener, router); err != nil {
		r.log.WithError(err).Error("failed to run server")
	}
}
