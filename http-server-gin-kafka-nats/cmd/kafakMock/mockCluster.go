package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	mockCluster, err := kafka.NewMockCluster(1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("mock cluster: %v", mockCluster.BootstrapServers())
	wg.Wait()
}
