package main

import (
	pb "gRPCnew/ClientSideStreaming/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedMyServiceServer
}

func (s *server) UploadData(stream pb.MyService_UploadDataServer) error {
	for {
		data, err := stream.Recv()
		if err != nil {
			// Handle end of stream or other errors
			return err
		}
		log.Printf("Received data: %s", data.Payload)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
