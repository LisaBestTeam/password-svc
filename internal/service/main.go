package service

import (
	"context"
	"sync"
)

type Service interface {
	Run(ctx context.Context, group *sync.WaitGroup)
}

func Runner(ctx context.Context, services ...Service) {
	group := &sync.WaitGroup{}
	for _, service := range services {
		group.Add(1)
		go service.Run(ctx, group)
	}

	group.Wait()
}
