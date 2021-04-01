package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	kafkaTopic   string
	kafkaGroup   string
	kafkaServers string
)

func init() {
	kafkaTopic = os.Getenv("KAFKA_TOPIC")
	kafkaGroup = os.Getenv("KAFKA_GROUP")
	kafkaServers = os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
}

func main() {
	ch := make(chan kafka.Message)

	go Consume(ch)

	for msg := range ch {
		log.Println("Recebendo mensagem:", string(msg.Value))
	}
}

func Consume(ch chan kafka.Message) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": kafkaServers,
		"group.id":          kafkaGroup,
	}

	consumer, err := kafka.NewConsumer(configMap)

	if err != nil {
		log.Fatal(err.Error())
	}

	consumer.Subscribe(kafkaTopic, nil)

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			ch <- *msg
		}
	}

}
