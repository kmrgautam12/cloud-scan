package service

import (
	"distributed-job-scheduler/jobqueue/internals/auth"
	"distributed-job-scheduler/jobqueue/internals/repository"
	"distributed-job-scheduler/jobqueue/utils/structure"
	"errors"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	region = "us-east-1"
)

const (
	secretString = "gtm-alx-supersecret-string"
	issuer       = "gtm-scheduler"
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

func (service *JobService) CreateJobs(req structure.CreateJobs) {

}

func (service *JobService) LoginUser(req structure.LoginUser) (userToken structure.UserAccessToken, err error) {
	exist, err := service.repo.CheckUserExistsByUsername(req.Username)
	if err != nil {
		return userToken, err
	}
	if !exist {
		return userToken, fmt.Errorf("user %s don't exist", req.Username)
	}
	user, err := service.repo.GetUserByUsername(req.Username)
	if err != nil {
		return userToken, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return userToken, errors.New("provided password is not valid")
	}
	token, err := auth.NewJwtService(secretString, issuer).GenerateToken(req.Username, "admin")
	if err != nil {
		return userToken, err
	}
	if token == "" {
		return userToken, errors.New("error generating token for customer")
	}
	tokenId, err := uuid.NewV4()
	if err != nil {
		return userToken, errors.New("error generating tokenId")
	}
	userToken = structure.UserAccessToken{
		TokenId: tokenId.String(),
		Token:   token,
	}
	return userToken, nil
}

func (service *JobService) CreateUser(body structure.CreateUser) error {
	hb, err := bcrypt.GenerateFromPassword([]byte(body.Password), 3)
	if err != nil {
		return err
	}
	passBody := body
	passBody.Password = string(hb)
	exist, err := service.repo.CheckUserExistsByUsername(passBody.UserName)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("username %s already exists", passBody.UserName)
	}
	err = service.repo.CreateUser(&passBody)
	if err != nil {
		return err
	}
	return nil
}

func (service *JobService) UploadFile(body structure.UploadFile) (string, error) {
	return service.repo.UploadFile(body)
}
