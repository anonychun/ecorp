package worker

import (
	"context"

	"github.com/go-co-op/gocron/v2"
)

func Start(ctx context.Context) error {
	cron, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	cron.Start()
	<-ctx.Done()

	return cron.Shutdown()
}
