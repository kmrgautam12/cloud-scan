package repository

import (
	"context"
	"distributed-job-scheduler/jobqueue/utils/structure"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jackc/pgx/v5/pgxpool"
)

type JobRepository struct {
	db  *pgxpool.Pool
	cfg cloudConfig
	c   context.Context
}

type cloudConfig struct {
	s3 *s3.Client
}

const (
	region = "us-east-1"
)

func NewJobRepository(db *pgxpool.Pool) (repo *JobRepository, err error) {
	cloudConfig, err := createAwsConfig()
	if err != nil {
		return repo, err
	}
	return &JobRepository{
		db:  db,
		cfg: cloudConfig,
		c:   context.Background(),
	}, nil
}

func createAwsConfig() (cloudCfg cloudConfig, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(lo *config.LoadOptions) error {
		lo.Region = region
		return nil
	})
	if err != nil {
		return cloudCfg, nil
	}
	s3Client := s3.NewFromConfig(cfg)
	if s3Client == nil {
		return cloudCfg, errors.New("failed creating cloud config")
	}
	cloudCfg.s3 = s3Client
	return cloudCfg, nil
}

func (h *JobRepository) CreateUser(body *structure.CreateUser) error {
	_, err := h.db.Exec(context.TODO(), createUser, body.UserName, body.Name, body.Password)
	if err != nil {
		return err
	}
	return nil
}

func (h *JobRepository) CheckUserExistsByUsername(username string) (bool, error) {
	var exist bool
	err := h.db.QueryRow(context.Background(), checkUserExistsByUsername, username).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *JobRepository) GetUserByUsername(username string) (user structure.CreateUser, err error) {
	err = h.db.QueryRow(context.Background(), getUserByUsername, username).Scan(&user.UserName, &user.Name, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (h *JobRepository) UploadFile(body structure.UploadFile) (string, error) {
	s3Client := h.cfg.s3
	presignClient := s3.NewPresignClient(s3Client)
	preSignResp, err := presignClient.PresignPutObject(h.c, &s3.PutObjectInput{
		Bucket: aws.String("gtm-test-bucket"),
		Key:    aws.String(fmt.Sprintf("scan/%s", body.FileName)),
	}, func(po *s3.PresignOptions) {
		po.Expires = time.Minute * 30
	})
	if err != nil {
		return "", err
	}
	url := preSignResp.URL
	if url == "" {
		return url, errors.New("empty presign url")
	}
	return url, nil
}

func CreatePresignedUrl()
