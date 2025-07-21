package producer

import (
    "context"
    "log"

    "github.com/segmentio/kafka-go"
    "kafka-go-rest/config"
)

func ProduceMessage(message string) error {
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers:  []string{config.KafkaBrokerAddress},
        Topic:    config.KafkaTopic,
        Balancer: &kafka.LeastBytes{},
    })
    defer writer.Close()

    err := writer.WriteMessages(context.Background(),
        kafka.Message{
            Key:   []byte("Key-A"),
            Value: []byte(message),
        },
    )
    if err != nil {
        log.Printf("Error writing message: %v", err)
        return err
    }
    log.Println("Message produced:", message)
    return nil
}
