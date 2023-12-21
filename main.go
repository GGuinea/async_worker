package main

import (
	"async_worker/config"
	"async_worker/internal"
	"async_worker/internal/backgroundJobs/workers"
	jobprocessor "async_worker/internal/pkg/jobProcessor"
	"async_worker/internal/pkg/jobProcessor/migrations"
	"context"
)

func main() {
	ctx := context.Background()

	cfg := config.NewConfig()
	compositionRoot := internal.NewCompositionRoot(ctx, cfg)
	backgroundJob, err := setupBackgroundJobProcessor(ctx, cfg, compositionRoot)
	if err != nil {
		panic(err)
	}
}

func setupBackgroundJobProcessor(ctx context.Context, cfg *config.Config, deps *internal.CompositionRoot) (*jobprocessor.JobProcessorClient, error) {
	err := migrations.PerformStartupRiverMigration(ctx, deps.DbPool)
	if err != nil {
		panic(err)
	}

	backgroundWorkersMgmnt := workers.NewBackgroundJobWorkers(deps.BackgroundWorkers)
	workers.AddDefaultWorker(backgroundWorkersMgmnt)
	workers.AddNewWorker(backgroundWorkersMgmnt, &workers.NewRequestWorker{})

	jobProcessorClient, err := jobprocessor.NewJobProcessorClient(ctx, deps, cfg.BackgroundProcessorConfig)
	if err != nil {
		return nil, err
	}

	return jobProcessorClient, nil
}
