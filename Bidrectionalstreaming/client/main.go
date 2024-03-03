package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "gRPCnew/Bidrectionalstreaming/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBidirectionalServiceClient(conn)

	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}

	// Send messages to the server
	for i := 0; i < 5; i++ {
		msg := &pb.Message{
			Text: "Hello from client",
		}
		if err := stream.Send(msg); err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
		time.Sleep(time.Second)
	}

	// Receive messages from the server
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive message: %v", err)
		}
		log.Printf("Received message from server: %s", msg.Text)
	}
}
