package main

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

const (
	topic        = "lesson24Topic"
	kafkaAddress = "localhost:9092"
)

func sendMsgToKafka(w *kafka.Writer, messages []string) {

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
	log.Println("Happy New Year...")
	fmt.Println("🎄")
}
