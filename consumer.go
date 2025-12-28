package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic        = "lesson24Topic"
	kafkaAddress = "localhost:9092"
	group        = "lesson24Group"
)

func readMsg(ctx context.Context, r *kafka.Reader) {

}
func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaAddress},
		Topic:   topic,
		GroupID: group,
	})
	defer r.Close()

	log.Println("Consumer connected...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	readMsg(ctx, r)
}
