package ingest

import (
	"github.com/lisabestteam/password-svc/internal/config"
	"github.com/lisabestteam/password-svc/internal/database"
	"github.com/lisabestteam/password-svc/internal/database/postgres"
	"github.com/lisabestteam/password-svc/internal/service"
	"gitlab.com/distributed_lab/logan/v3"
)

type ingest struct {
	log     *logan.Entry
	channel <-chan database.Password
	db      database.Passwords
}

func NewIngest(cfg config.Config, channel <-chan database.Password) service.Service {
	return &ingest{
		log:     cfg.Log(),
		db:      postgres.NewPassword(cfg.DB()),
		channel: channel,
	}
}
