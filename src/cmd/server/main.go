package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	pb "alien-contact-detector/internal/pkg"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type TransmitterServer struct {
	pb.UnimplementedTransmitterServiceServer
	logFile *os.File // указатель на файл логирования
}

// Инициализация файла логов
func newTransmitterServer(logPath string) *TransmitterServer {
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	return &TransmitterServer{logFile: file}
}

func GetSessionId() string {
	uuid := uuid.New().String()
	// log.Println("Generated uuid: ", uuid)

	return uuid
}

func GetRandomFrequency() float64 {
	mean := rand.Float64()*20 - 10           // по заданию [-10, 10]
	stdDev := rand.Float64()*(1.5-0.3) + 0.3 // [0.3, 1.5]

	frequency := mean + stdDev*rand.NormFloat64()

	// log.Println("Generated random mean: ", mean)
	// log.Println("Generated random standard deviation: ", stdDev)
	// log.Println("Got frequency: ", frequency)

	return frequency
}

func GetTimestamp() int64 {
	timestamp := time.Now().Unix()
	// log.Println("Generated timestamp: ", timestamp)

	return timestamp
}

func (s *TransmitterServer) GenerateFrequency(req *pb.FrequencyRequest, stream pb.TransmitterService_GenerateFrequencyServer) error {
	sessionID := GetSessionId()
	s.logFile.WriteString(fmt.Sprintf("Session %s started \n", sessionID))
	for i := 0; i < int(req.GetNumValues()); i++ {

		freqData := &pb.Frequency{
			SessionId: sessionID,
			Frequency: GetRandomFrequency(),
			Timestamp: GetTimestamp(),
		}
		// Логируем каждую запись
		logEntry := fmt.Sprintf("SessionID: %s | Frequency: %.2f | Timestamp: %d\n", freqData.SessionId, freqData.Frequency, freqData.Timestamp)
		log.Print(logEntry)
		if _, err := s.logFile.WriteString(logEntry); err != nil {
			return fmt.Errorf("failed to write log to file: %v", err)
		}
		if err := stream.Send(freqData); err != nil {
			return err
		}
		time.Sleep(time.Second) // Симуляция интервала отправки данных
	}
	endLogEntry := fmt.Sprintf("Session %s ended\n", sessionID)
	log.Print(endLogEntry)
	if _, err := s.logFile.WriteString(endLogEntry); err != nil {
		return fmt.Errorf("failed to write log to file: %v", err)
	}
	return nil
}

func main() {
	logPath := "logs/server.log"
	server := newTransmitterServer(logPath)
	defer server.logFile.Close()
	lis, err := net.Listen("tcp", "localhost:3333")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTransmitterServiceServer(grpcServer, server)
	log.Printf("Server listening on localhost:3333")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
