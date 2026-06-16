package service

import (
	"distributed-job-scheduler/jobqueue/internals/repository"
	"distributed-job-scheduler/jobqueue/utils/structure"
)

type JobService struct {
	repo *repository.JobRepository
}

type SchedulerService interface {
	CreateJobs(req structure.CreateJobs) error
}

func NewJobService(repo *repository.JobRepository) *JobService {
	return &JobService{
		repo: repo,
	}
}
