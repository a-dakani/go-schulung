package kafka

import (
	"encoding/json"
	"github.com/a-dakani/go-schulung/http-server-gin-kafka/ginserver"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"golang.org/x/net/context"
)

type AutoNotifier struct {
	p *kafka.Producer
}

func NewAutoNotifier() *AutoNotifier {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "127.0.0.1:60516"})
	if err != nil {
		panic(err)
	}
	return &AutoNotifier{p: p}
}

func (an *AutoNotifier) NewAutoCreated(ctx context.Context, auto ginserver.Auto) error {
	//	send a kafka message

	autoJSON, err := json.Marshal(auto)
	if err != nil {
		return err
	}
	dv := make(chan kafka.Event)
	topic := "new-auto-created"

	err = an.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          autoJSON,
	}, dv)

	return nil
}
