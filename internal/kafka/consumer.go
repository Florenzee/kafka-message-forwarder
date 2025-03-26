package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

// Membaca message dari Kafka dan mengirim ke channel
func Consumer(brokers []string, sourceTopic, groupID string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    sourceTopic,
		GroupID:  groupID,
		MaxBytes: 10e6, // 10MB per pesan
	})

	// Loop membaca pesan dari Kafka
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading Kafka message: %v", err)
			continue
		}

		// Format message sebagai JSON
		var formattedJSON map[string]interface{}
		if err := json.Unmarshal(msg.Value, &formattedJSON); err != nil {
			log.Printf("Invalid JSON message: %s", string(msg.Value))
			continue
		}

		// Kirim pesan ke channel
		MessageChannel <- msg
	}
}
