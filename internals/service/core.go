package service

import "distributed-job-scheduler/jobqueue/utils/structure"

type Service struct {
	db string
}

type SchedulerService interface {
	CreateJobs(req structure.CreateJobs) error
}

func NewService() Service {
	return Service{}
}
func (s *Service) CreateJobs() error {
	// return nil
	return nil
}
