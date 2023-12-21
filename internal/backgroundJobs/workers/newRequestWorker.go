package workers

import (
	"context"
	"fmt"
	"time"

	"github.com/riverqueue/river"
)

type NewRequestArgs struct {
	User string
}

func (NewRequestArgs) Kind() string { return "new_request" }

type NewRequestWorker struct {
	river.WorkerDefaults[NewRequestArgs]
}

func (w *NewRequestWorker) Work(ctx context.Context, job *river.Job[NewRequestArgs]) error {
	time.Sleep(5 * time.Second)
	fmt.Printf("Starting processing newRequestJob for: %s\n", job.Args.User)
	for i := 0; i < 5; i++ {
		fmt.Printf("Processing newRequestJob for: %s\n", job.Args.User)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Finished processing newRequestJob")
	return nil
}
