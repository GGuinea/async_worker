package internal

import (
	"async_worker/config"
	"context"

	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
)

type CompositionRoot struct {
	BackgroundWorkers *river.Workers
	DbPool            *pgxpool.Pool
}

func NewCompositionRoot(ctx context.Context, config *config.Config) *CompositionRoot {
	dbPool, err := getDbPool(ctx, &config.Db)
	if err != nil {
		panic(err)
	}

	backgroundWorkers, err := initBackgroundJobWorkers()
	if err != nil {
		panic(err)
	}

	return &CompositionRoot{DbPool: dbPool, BackgroundWorkers: backgroundWorkers}
}

func initBackgroundJobWorkers() (*river.Workers, error) {
	workers := river.NewWorkers()
	if workers == nil {
		return nil, fmt.Errorf("failed to create workers")
	}

	return workers, nil
}

func getDbPool(ctx context.Context, config *config.DbConfig) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, buildConnectionString(config))
}

func buildConnectionString(dbConfig *config.DbConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
}
