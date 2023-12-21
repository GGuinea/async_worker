package app

import (
	"async_worker/internal/backgroundJobs/workers"
	jobprocessor "async_worker/internal/pkg/jobProcessor"
	"context"
)

type Service struct {
	JobProcessor *jobprocessor.JobProcessorClient
}

func NewService(jobProcessor *jobprocessor.JobProcessorClient) *Service {
	return &Service{
		JobProcessor: jobProcessor,
	}
}

func (s *Service) HandleServiceRequest(ctx context.Context, user string) error {
	err := s.JobProcessor.ScheduleNewJob(ctx, workers.NewRequestArgs{
		User: user,
	})
	if err != nil {
		return err
	}

	return nil
}
