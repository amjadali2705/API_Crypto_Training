package consumer

import (
	"context"
	"log"

	"kafka-go-rest/config"

	"github.com/segmentio/kafka-go"
)

func StartConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.KafkaBrokerAddress},
		Topic:   config.KafkaTopic,
		GroupID: "group-id",
	})

	go func() {
		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}
			log.Printf("Received message: %s = %s", string(msg.Key), string(msg.Value))
		}
	}()
}
