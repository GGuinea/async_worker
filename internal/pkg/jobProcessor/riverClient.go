package jobprocessor

import (
	"async_worker/config"
	"async_worker/internal"
	"context"

	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
)

type JobProcessorClient struct {
	RiverClient *river.Client[pgx.Tx]
	DbPool      *pgxpool.Pool
}

func NewJobProcessorClient(ctx context.Context, deps *internal.CompositionRoot, cfg config.BackgroundProcessorConfig) (*JobProcessorClient, error) {
	riverClient, err := river.NewClient(riverpgxv5.New(deps.DbPool), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: cfg.MaxWorkers},
		},
		Workers: deps.BackgroundWorkers,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	if err := riverClient.Start(ctx); err != nil {
		return nil, fmt.Errorf("failed to start client: %w", err)
	}

	return &JobProcessorClient{RiverClient: riverClient, DbPool: deps.DbPool}, nil
}

func (jbc *JobProcessorClient) ScheduleNewJob(ctx context.Context, args river.JobArgs) error {
	tx, err := jbc.DbPool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin tx: %w", err)
	}

	_, err = jbc.RiverClient.InsertTx(ctx, tx, args, nil)
	if err != nil {
		return fmt.Errorf("failed to insert new job: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit tx: %w", err)
	}
	return nil
}
