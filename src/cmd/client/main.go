package main

import (
	"context"
	"log"

	pb "alien-contact-detector/internal/pkg" // убедитесь, что путь правильный

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:3333", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTransmitterServiceClient(conn)

	// Это хорошая практика
	numValues := int32(1) // количество запросов

	stream, err := client.GenerateFrequency(context.Background(), &pb.FrequencyRequest{NumValues: numValues})
	if err != nil {
		log.Fatalf("Error calling GenerateFrequency: %v", err)
	}

	for {
		data, err := stream.Recv()
		if err != nil {
			log.Printf("Error receiving data: %v", err)
			break
		}
		log.Printf("Received data: session_id=%s, frequency=%.2f, timestamp=%d", data.SessionId, data.Frequency, data.Timestamp)
	}
}
