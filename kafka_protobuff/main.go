package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	elliot := &Person{
		Name: "Elliot",
		Age:  25,
	}
	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(data)
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		panic(err)
	}

	defer p.Close()
	topic := "test-consumer"

	deliveryChan := make(chan kafka.Event)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, deliveryChan)

	select {
	case ev := <-deliveryChan:
		switch e := ev.(type) {
		case *kafka.Message:
			if e.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", e.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %v\n", e.TopicPartition)
				fmt.Printf("Delivered message to %v\n", string(e.Value))
			}
			close(deliveryChan)
		}
	}
}
