package scanserver

import (
	"context"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ScannerService struct {
	ctx      context.Context
	cloudCfg *cloudConfig
}

type cloudConfig struct {
	s3 *s3.Client
}

const region = "us-east-1"

func NewScannerService() *ScannerService {
	return &ScannerService{
		ctx:      context.Background(),
		cloudCfg: NewCloudConfig(),
	}
}

func NewCloudConfig() (cloudCfg *cloudConfig) {
	cfg, err := config.LoadDefaultConfig(context.Background(), func(lo *config.LoadOptions) error {
		lo.Region = region
		lo.RetryMaxAttempts = 3
		lo.RetryMode = aws.RetryModeStandard
		return nil
	},
	)
	if err != nil {
		log.Printf("Error creating cloud config, err %v", err)
		panic(err)
	}

	s3Client := s3.NewFromConfig(cfg)
	if s3Client == nil {
		log.Printf("Error creating s3 client from config, error: %v", err)
		panic("Unable to create s3 client config")
	}
	cloudCfg.s3 = s3Client
	return cloudCfg
}

func (s *ScannerService) ScanFile(jobId, bucket, key string) (infected bool, msg string, err error) {
	object, err := s.cloudCfg.s3.GetObject(s.ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
	},
	)
	if err != nil {
		return false, "", err
	}
	b, err := io.ReadAll(object.Body)
	if err != nil {
		return false, "", err
	}
}
