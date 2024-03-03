package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "gRPCnew/Bidrectionalstreaming/pb"
)

type server struct{
	pb.UnimplementedBidirectionalServiceServer
}

func (s *server) Chat(stream pb.BidirectionalService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Received message from client: %s", msg.Text)

		// Echo the message back to the client
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterBidirectionalServiceServer(srv, &server{})

	log.Println("Server started at :50051")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
