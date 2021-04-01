package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	kafkaTopic   string
	kafkaServers string
)

func init() {
	kafkaTopic = os.Getenv("KAFKA_TOPIC")
	kafkaServers = os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
}

func main() {

	forever := make(chan bool)

	producer := NewKafkaProducer()

	Publish("Olá Mundo GO-Kafka", kafkaTopic, producer)

	<-forever
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": kafkaServers,
	}
	p, err := kafka.NewProducer(configMap)

	if err != nil {
		log.Fatal(err.Error())
	}
	return p
}

func Publish(msg, topic string, producer *kafka.Producer) error {
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := producer.Produce(message, nil)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
