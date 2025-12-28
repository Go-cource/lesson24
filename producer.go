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
)

func sendMsgToKafka(w *kafka.Writer, messages []string) {
	for i, msg := range messages {
		oneMessage := kafka.Message{
			Key:   []byte(fmt.Sprintf("Message #%d", i+1)),
			Value: []byte(msg),
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := w.WriteMessages(ctx, oneMessage)
		if err != nil {
			fmt.Println("Error with kafka writer: ", err)
		} else {
			fmt.Printf("Message #%d sended to kafka\n", i+1)
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{kafkaAddress},
		Topic:   topic,
	})
	defer w.Close()
	messages := []string{
		"Last Christmas",
		"I gave u my heart",
		"But in very next day",
		"U gave it away",
	}
	sendMsgToKafka(w, messages)
	log.Println("Producer finished...")
	log.Println("Happy New Year!")
	fmt.Println("🎄")
}
