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

func GetRandomFrequency() float64 {
	mean := rand.Float64()*20 - 10           // по заданию [-10, 10]
	stdDev := rand.Float64()*(1.5-0.3) + 0.3 // [0.3, 1.5]

	frequency := mean + stdDev*rand.NormFloat64()

	log.Println("Generated random mean: ", mean)
	log.Println("Generated random standard deviation: ", stdDev)
	log.Println("Got frequency: ", frequency)

	return frequency
}

func GetTimestamp() int64 {
	timestamp := time.Now().Unix()
	log.Println("Generated timestamp: ", timestamp)

	return timestamp
}

func (s *TransmitterServer) GenerateFrequency(req *pb.FrequencyRequest, stream pb.TransmitterService_GenerateFrequencyServer) error {
	for i := 0; i < int(req.GetNumValues()); i++ {

		freqData := &pb.Frequency{
			SessionId: GetSessionId(),
			Frequency: GetRandomFrequency(),
			Timestamp: GetTimestamp(),
		}

		if err := stream.Send(freqData); err != nil {
			return err
		}
		time.Sleep(time.Second) // Симуляция интервала отправки данных
	}
	return nil
}

func main() {
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
