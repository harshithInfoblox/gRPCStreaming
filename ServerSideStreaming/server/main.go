package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "gRPCnew/ServerSideStreaming/pb"
	"google.golang.org/grpc"
)

// Server implements the StreamServiceServer interface.
type Server struct{
	pb.UnimplementedStreamServiceServer
}

// FetchResponse handles FetchResponse RPC.
func (s *Server) FetchResponse(req *pb.Request, srv pb.StreamService_FetchResponseServer) error {
	log.Printf("Received request for ID: %d", req.Id)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(count int) {
			defer wg.Done()

			// Simulate processing time
			time.Sleep(time.Duration(count) * time.Second)
			resp := pb.Response{Result: fmt.Sprintf("Response #%d for ID:%d", count+1, req.Id)}
			if err := srv.Send(&resp); err != nil {
				log.Printf("Failed to send response: %v", err)
			}
			log.Printf("Finished processing response #%d for ID: %d", count+1, req.Id)
		}(i)
	}

	wg.Wait()
	return nil
}

func main() {
	// Create listener
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, &Server{})

	log.Println("Starting server")
	// Start serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
