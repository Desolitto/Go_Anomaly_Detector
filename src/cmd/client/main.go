package main

import (
	"context"
	"log"

	pb "alien-contact-detector/internal/pkg" // убедитесь, что путь правильный

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Установление соединения с сервером gRPC
	conn, err := grpc.NewClient("localhost:3333", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Создание клиента для TransmitterService
	client := pb.NewTransmitterServiceClient(conn)

	// Установка количества значений для генерации
	numValues := int32(1) // измените это значение по необходимости

	// Создание запроса для получения данных частоты
	stream, err := client.GenerateFrequency(context.Background(), &pb.FrequencyRequest{NumValues: numValues})
	if err != nil {
		log.Fatalf("Error calling GenerateFrequency: %v", err)
	}

	// Обработка полученных данных в потоке
	for {
		data, err := stream.Recv()
		if err != nil {
			log.Printf("Error receiving data: %v", err)
			break // Завершение цикла при ошибке
		}
		log.Printf("Received data: session_id=%s, frequency=%.2f, timestamp=%d", data.SessionId, data.Frequency, data.Timestamp)
	}
}
