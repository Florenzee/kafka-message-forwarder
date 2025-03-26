package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// Meneruskan message dari channel ke destination
func Forwarder(brokers []string, destinationTopic string) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    destinationTopic,
		Async:    false,
		Balancer: &kafka.LeastBytes{},
	})

	for msg := range MessageChannel {
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Key:   msg.Key,
			Value: msg.Value,
		})
		if err != nil {
			log.Printf("Error forwarding message: %v", err)
		} else {
			log.Printf("Message successfully forwarded to topic: %s", destinationTopic)
		}
	}
}
