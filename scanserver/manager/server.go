package scanserver

import (
	"context"
	"distributed-job-scheduler/jobqueue/stub/proto"

	"google.golang.org/grpc/status"
)

type ScanServer struct {
	proto.UnimplementedScannerServiceServer
	service *ScannerService
}

func NewScanServer(service *ScannerService) *ScanServer {
	return &ScanServer{
		service: service,
	}
}

func (s *ScanServer) ScanFile(ctx context.Context, req *proto.ScanRequest) (res *proto.ScanResponse, err error) {
	if req.JobId == "" || req.Bucket == "" || req.ObjectKey == "" {
		return nil, status.Error(400, "Invalid request body")
	}
	infected, msg := s.service.ScanFile(req.JobId, req.Bucket, req.ObjectKey)
	res = &proto.ScanResponse{
		Infected: infected,
		Message:  msg,
	}
	return res, nil
}
