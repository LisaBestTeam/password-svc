package ingest

import (
	"context"
	"sync"
)

func (i *ingest) Run(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()

	log := i.log.WithField("service", "ingest")
	log.Info("run ingest")

	select {
	case <-ctx.Done():
		log.Info("close ingest")
		return
	case password := <-i.channel:
		if err := i.db.CreatePassword(password); err != nil {
			log.WithError(err).Error("failed to create password in database")
		}
	}
}
