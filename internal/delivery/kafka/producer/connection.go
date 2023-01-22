package producer

import (
	"net"
	"os"

	"github.com/segmentio/kafka-go"
)

var (
	kafkaHost = os.Getenv("KAFKA_HOST")
	kafkaPort = os.Getenv("KAFKA_PORT")
)

func Connection() *kafka.Conn {
	connection, err := kafka.Dial("tcp", net.JoinHostPort(kafkaHost, kafkaPort))

	if err != nil {
		panic("kafka not ready")
	}

	return connection
}
