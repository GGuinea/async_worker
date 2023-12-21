package migrations

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	"github.com/riverqueue/river/rivermigrate"
)

func PerformStartupRiverMigration(ctx context.Context, dbPool *pgxpool.Pool) error {
	err := migrate(ctx, dbPool)

	if err != nil {
		return fmt.Errorf("Error performin startup river migration: %w", err)
	}

	return nil
}

func migrate(ctx context.Context, dbPool *pgxpool.Pool) error {
	migrator := rivermigrate.New(riverpgxv5.New(dbPool), nil)

	tx, err := dbPool.Begin(ctx)
	if err != nil {
		panic(err)
	}

	_, err = migrator.MigrateTx(ctx, tx, rivermigrate.DirectionUp, &rivermigrate.MigrateOpts{})

	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
