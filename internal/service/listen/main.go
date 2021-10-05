package listen

import (
	"github.com/lisabestteam/password-svc/internal/config"
	"github.com/lisabestteam/password-svc/internal/database"
	"github.com/lisabestteam/password-svc/internal/database/postgres"
	"github.com/lisabestteam/password-svc/internal/horizon"
	"github.com/lisabestteam/password-svc/internal/service"
	"github.com/sirupsen/logrus"
)

type listen struct {
	passwords database.Passwords
	client    horizon.HorizonClient
	log       *logrus.Logger
	channel   chan<- database.Password
}

func NewListen(cfg config.Config, channel chan<- database.Password) service.Service {
	return &listen{
		passwords: postgres.NewPassword(cfg.Database()),
		client:    cfg.Horizon(),
		log:       cfg.Log(),
		channel:   channel,
	}
}
