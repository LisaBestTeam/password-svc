package cli

import (
	"context"

	"github.com/lisabestteam/password-svc/internal/config"
	"github.com/lisabestteam/password-svc/internal/database"
	service2 "github.com/lisabestteam/password-svc/internal/service"
	ingest2 "github.com/lisabestteam/password-svc/internal/service/ingest"
	"github.com/lisabestteam/password-svc/internal/service/listen"
	"github.com/lisabestteam/password-svc/internal/service/router"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	application = kingpin.New("password", "a command line")

	configPath = application.Flag("config", "path to config file").Required().ExistingFile()

	migration     = application.Command("migrate", "migrate command")
	migrationUp   = migration.Command("up", "apply migrations")
	migrationDown = migration.Command("down", "cancel migrations")

	running = application.Command("run", "")
	service = running.Command("service", "")
	ingest  = service.Command("ingest", "")
	server  = service.Command("server", "")
)

func Run(args []string) bool {
	cmd := kingpin.MustParse(application.Parse(args[1:]))

	cfg := config.NewConfig(config.MustGetter(*configPath))
	ctx := context.Background()
	log := cfg.Log()

	defer cfg.Database().Close()

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
	case server.FullCommand():
		service2.Runner(ctx, router.NewRouter(cfg))
	}

	<-ctx.Done()

	return true
}
