package service

import (
	"context"
	"distributed-job-scheduler/jobqueue/internals/repository"
	"distributed-job-scheduler/jobqueue/utils/structure"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	region = "us-east-1"
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
func ConnectCloud() {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		func(lo *config.LoadOptions) error {
			lo.Region = region
			lo.Retryer = aws.NewConfig().Retryer
			return nil
		})
	if err != nil {

	}
}
func (service *JobService) CreateJobs(req structure.CreateJobs) {

}
