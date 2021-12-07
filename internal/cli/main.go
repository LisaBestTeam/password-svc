package cli

import (
	"context"

	"github.com/lisabestteam/password-svc/internal/config"
	"github.com/lisabestteam/password-svc/internal/database"
	service2 "github.com/lisabestteam/password-svc/internal/service"
	ingest2 "github.com/lisabestteam/password-svc/internal/service/ingest"
	"github.com/lisabestteam/password-svc/internal/service/listen"
	"github.com/lisabestteam/password-svc/internal/service/router"
	"github.com/sirupsen/logrus"
	"gitlab.com/distributed_lab/kit/kv"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	application = kingpin.New("password", "a command line")

	migration     = application.Command("migrate", "migrate command")
	migrationUp   = migration.Command("up", "apply migrations")
	migrationDown = migration.Command("down", "cancel migrations")

	running = application.Command("run", "command for run")
	service = running.Command("service", "command for choose service")
	ingest  = service.Command("ingest", "ingest passwords from blockchain")
	server  = service.Command("server", "http server passwords")
)

func Run(args []string) bool {
	defer func() {
		if msg := recover(); msg != nil {
			logrus.New().WithField("reason", msg).Fatal("app panicked")
		}
	}()

	cmd := kingpin.MustParse(application.Parse(args[1:]))

	cfg := config.NewConfig(kv.MustFromEnv())
	ctx := context.Background()
	log := cfg.Log()

	switch cmd {
	case migrationUp.FullCommand():
		if err := MigrateUp(cfg); err != nil {
			log.WithError(err).Error("failed migrate")
		}
		return true
	case migrationDown.FullCommand():
		if err := MigrateDown(cfg); err != nil {
			log.WithError(err).Error("failed migrate")
		}
		return true
	case ingest.FullCommand():
		channel := make(chan database.Password, 1)
		service2.Runner(ctx, ingest2.NewIngest(cfg, channel), listen.NewListen(cfg, channel))
		close(channel)
	case server.FullCommand():
		service2.Runner(ctx, router.NewServer(cfg))
	}

	return true
}
