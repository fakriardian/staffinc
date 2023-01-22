package consumer

import (
	"os"

	"github.com/segmentio/kafka-go"
)

var (
	kafkaURL = os.Getenv("KAFKA_URL")
)

func GetKafkaReader(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}
