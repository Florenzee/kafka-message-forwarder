package kafka

import "github.com/segmentio/kafka-go"

// Mmenghubungkan Consumer dan Forwarder
var MessageChannel = make(chan kafka.Message, 100)
