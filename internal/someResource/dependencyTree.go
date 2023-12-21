package someresource

import jobprocessor "async_worker/internal/pkg/jobProcessor"

type SomeResourceDependencies struct {
	JobProcessor *jobprocessor.JobProcessorClient
}

func NewSomeResourceDependencies(jobProcessor *jobprocessor.JobProcessorClient) *SomeResourceDependencies {
	if jobProcessor == nil {
		panic("jobProcessor is nil")
	}

	return &SomeResourceDependencies{
		JobProcessor: jobProcessor,
	}
}
