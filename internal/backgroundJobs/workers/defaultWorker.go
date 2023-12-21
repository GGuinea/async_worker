package workers

import (
	"context"

	"github.com/riverqueue/river"
)

type DefaultArgs struct {
}

func (DefaultArgs) Kind() string { return "default" }

type DefautlWorker struct {
	river.WorkerDefaults[DefaultArgs]
}

func (w *DefautlWorker) Work(ctx context.Context, job *river.Job[DefaultArgs]) error {
	return nil
}
