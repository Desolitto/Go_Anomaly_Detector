package main

import (
	"context"
	"flag"
	"io"
	"log"
	"math"
	"os"

	pb "alien-contact-detector/internal/pkg" // убедитесь, что путь правильный

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	defaultK       = 1.5
	defaultAddress = "localhost:3333"
)

func parseFlags() (k *float64) {
	k = flag.Float64("k", defaultK, "Anomaly detection coefficient")
	flag.Parse()
	return
}

func setupLogging() *os.File {
	// Создаем или открываем файл для логов
	file, err := os.OpenFile("logs/client.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, file)
	// Настраиваем логирование
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // Формат логов

	return file
}

func analyzeFrequencies(stream pb.TransmitterService_GenerateFrequencyClient, k float64) error {
	var (
		sum    float64 // Сумма всех частот для вычисления среднего значения
		sumSq  float64 // Сумма квадратов частот для вычисления дисперсии
		count  int     // Счетчик полученных сообщений
		mean   float64 // Среднее значение частот
		stdDev float64 // Стандартное отклонение частот
	)
	for {
		data, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("Analysis completed.")
				return nil // Exit the function without error
			}
			return status.Convert(err).Err() // Return the error for handling in main
		}
		count++
		frequency := data.GetFrequency()
		sum += frequency
		sumSq += frequency * frequency
		if count > 0 {
			mean = sum / float64(count)
			variance := (sumSq / float64(count)) - (mean * mean) // дисперсии
			if variance < 0 {
				variance = 0
			}
			stdDev = math.Sqrt(variance)
		}
		log.Printf("Received frequency: %f, Mean: %f, StdDev: %f", frequency, mean, stdDev)
		if math.Abs(frequency-mean) > k*stdDev {
			log.Printf("Anomaly detected! Frequency: %f, Mean: %f, StdDev: %f", frequency, mean, stdDev)
		}
	}
}

func main() {

	k := parseFlags()
	logFile := setupLogging()
	defer logFile.Close()
	// log.Printf("Anomaly detection coefficient: %f", *k)
	conn, err := grpc.NewClient(defaultAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTransmitterServiceClient(conn)

	// Это хорошая практика
	numValues := int32(2) // количество запросов

	stream, err := client.GenerateFrequency(context.Background(), &pb.FrequencyRequest{NumValues: numValues})
	if err != nil {
		log.Fatalf("Error calling GenerateFrequency: %v", err)
	}

	if err := analyzeFrequencies(stream, *k); err != nil {
		log.Printf("Error during analysis: %v", err)
	}
}
