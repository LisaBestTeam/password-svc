package listen

import (
	"github.com/lisabestteam/password-svc/internal/config"
	"github.com/lisabestteam/password-svc/internal/database"
	"github.com/lisabestteam/password-svc/internal/database/postgres"
	"github.com/lisabestteam/password-svc/internal/horizon"
	"github.com/lisabestteam/password-svc/internal/service"
	"gitlab.com/distributed_lab/logan/v3"
)

type listen struct {
	passwords database.Passwords
	client    horizon.HorizonClient
	log       *logan.Entry
	db        database.Passwords
}

func NewListen(cfg config.Config) service.Service {
	return &listen{
		passwords: postgres.NewPassword(cfg.DB()),
		client:    horizon.NewHorizonClient("http://horizon"),
		log:       cfg.Log(),
		db:        postgres.NewPassword(cfg.DB()),
	}
}
