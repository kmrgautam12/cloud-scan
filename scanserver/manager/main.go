package scanserver

import (
	"distributed-job-scheduler/jobqueue/stub/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	// Register the scanner server scan service
	scannerService := NewScannerService()
	scanServer := NewScanServer(scannerService)

	// Build grpc server
	grpc := grpc.NewServer()
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Printf("error creating rpc server %v", err)
		panic(err)
	}
	// Register Services
	proto.RegisterScannerServiceServer(grpc, scanServer)
	log.Println("Scanner gRPC server on :50051 port")

	// Listen the grpc server
	if err := grpc.Serve(listener); err != nil {
		log.Printf("error serving gRPC; error %v", err)
		panic(err)
	}

}
