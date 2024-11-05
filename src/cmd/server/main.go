package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	pb "alien-contact-detector/internal/pkg"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type TransmitterServer struct {
	pb.UnimplementedTransmitterServiceServer
}

func GetSessionId() string {
	uuid := uuid.New().String()
	log.Println("Generated uuid: ", uuid)

	return uuid
}

func (s *TransmitterServer) GenerateFrequency(req *pb.FrequencyRequest, stream pb.TransmitterService_GenerateFrequencyServer) error {
	for i := 0; i < int(req.GetNumValues()); i++ {
		frequency := rand.Float64() * 100.0
		timestamp := time.Now().Unix()

		freqData := &pb.Frequency{
			SessionId: GetSessionId(),
			Frequency: frequency,
			Timestamp: timestamp,
		}

		// Здесь правильно отправляется *pkg.Frequency
		if err := stream.Send(freqData); err != nil {
			return err
		}
		time.Sleep(time.Second) // Симуляция интервала отправки данных
	}
	return nil
}

func main() {
	// Создаём gRPC-сервер
	lis, err := net.Listen("tcp", "localhost:3333")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTransmitterServiceServer(grpcServer, &TransmitterServer{})
	log.Printf("Server listening on localhost:3333")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
