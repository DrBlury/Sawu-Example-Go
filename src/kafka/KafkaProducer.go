package kafka

import (
	"fmt"
	"os"
	"sawu-example-go/config"
	"sawu-example-go/entities"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

var p *kafka.Producer

// CreateProducer creates the Kafka Producer
func CreateProducer() {
	//Set default broker ip if not set
	broker, isPresent := os.LookupEnv("kafka_broker_ip")
	if isPresent == false {
		broker = config.Defaults.Kafka.Broker.IPAddress
	}
	fmt.Println(broker)
	var err error
	p, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created Producer %v\n", p)
}

// SendNextStepEvent sends a NextStepEvent to Kafka
func SendNextStepEvent(nextStepEvent entities.NextStepEvent) {
	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)
	nextStepEvent.Internal = "true"
	// Set topic name
	topic := fmt.Sprintf("%s-%s", nextStepEvent.ProcessName, nextStepEvent.ProcessStep)
	log.Info(fmt.Sprintf("Sending Event to topic: %s", topic))
	value := serialize(nextStepEvent)
	log.Info(value)
	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, deliveryChan)

	if err != nil {
		log.Error(err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}

func serialize(nextStepEvent entities.NextStepEvent) string {
	event := fmt.Sprintf("id=%s,timestamp=%s,processName=%s,comingFromID=%s,processInstanceID=%s,processStep=%s,internal=%s,retryCount=%s,$e%%,%s",
		nextStepEvent.ID, nextStepEvent.TimeStamp, nextStepEvent.ProcessName, nextStepEvent.ComingFromID, nextStepEvent.ProcessInstanceID, nextStepEvent.ProcessStep, nextStepEvent.Internal, nextStepEvent.RetryCount, nextStepEvent.Data)

	return event
}
