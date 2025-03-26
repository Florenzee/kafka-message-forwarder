package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bitbucket.org/Amartha/go-dlq-retrier/internal/handler"
	"bitbucket.org/Amartha/go-dlq-retrier/internal/handler/health"
	server "bitbucket.org/Amartha/go-dlq-retrier/internal/http"
	"bitbucket.org/Amartha/go-dlq-retrier/internal/kafka"

	"github.com/joho/godotenv"
)

func main() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Ambil data dari environment
	broker := os.Getenv("brokerKafka")
	source := os.Getenv("sourceTopic")
	ID := os.Getenv("groupID")
	destination := os.Getenv("destinationTopic")

	// Shutdown signal
	appContext, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer cancel()

	// Start HTTP server
	srv := server.New(handler.New(health.NewGet()))
	srv.Start(appContext)

	// Start Kafka Consumers dalam goroutine
	go kafka.Consumer(
		[]string{broker}, // Kafka Broker
		source,           // Source Topic
		ID,               // Group ID untuk Consumer 1
	)

	go kafka.Forwarder(
		[]string{broker}, // Kafka Broker
		destination,      // Destination Topic
	)

	// Tunggu sinyal shutdown
	<-appContext.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Shutdown HTTP server
	srv.Close(ctx)
}
