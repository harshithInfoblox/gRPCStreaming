package main

import (
	"context"
	pb "gRPCnew/ClientSideStreaming/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMyServiceClient(conn)

	stream, err := client.UploadData(context.Background())
	if err != nil {
		log.Fatalf("could not upload data: %v", err)
	}

	// Send multiple data payloads
	payloads := []string{"payload1", "payload2", "payload3"}
	for _, payload := range payloads {
		if err := stream.Send(&pb.Data{Payload: payload}); err != nil {
			log.Fatalf("error sending data: %v", err)
		}
	}

	// Close and receive the response from the server
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response: %v", err)
	}
	log.Printf("Server response: %s", response.Message)
}
