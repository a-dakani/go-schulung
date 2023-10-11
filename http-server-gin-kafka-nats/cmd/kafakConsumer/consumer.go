package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:60516",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{"new-auto-created"}, nil)
	if err != nil {
		panic(err)
	}
	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(msg.Value))
	}

}
