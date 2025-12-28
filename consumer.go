package main

import (
	"context"
	"fmt"
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
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("I got message: %s:%s (topic: %v, partition: %v, offset: %v)\n", string(msg.Key), string(msg.Value), msg.Topic, msg.Partition, msg.Offset)

	}
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
