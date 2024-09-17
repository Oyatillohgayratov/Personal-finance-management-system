package kafkaa

import (
	"encoding/json"
	"log"
	"user-service/protos/notifactionpb"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// func ConnectKafka(kurl string) (*Kafka, error) {
// 	client, err := kgo.NewClient(
// 		kgo.SeedBrokers(kurl),
// 		kgo.AllowAutoTopicCreation(),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed kafka Client error: %v", err)
// 	}

// 	err = client.Ping(context.Background())
// 	if err != nil {
// 		return nil, fmt.Errorf("failed kafka connection: %v", err)
// 	}
// 	return &Kafka{Client: client}, nil
// }

func ProduceRegistrationEmail(email, sub, mes string) error {
	broker := "localhost:9092"
	topic := "notifications"

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer p.Close()

	emailMsg := &notifactionpb.NotificationReq{
		UserEmail: email,
		Subject:   sub,
		Message:   mes,
	}

	msgBytes, err := json.Marshal(emailMsg)
	if err != nil {
		log.Fatalf("Failed to marshal email message: %s", err)
	}

	deliveryChan := make(chan kafka.Event)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msgBytes,
	}, deliveryChan)
	if err != nil {
		log.Fatalf("Failed to produce message: %s", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		log.Fatalf("Delivery failed: %v", m.TopicPartition.Error)
	} else {
		log.Printf("Delivered message to %v", m.TopicPartition)
	}
	return nil
}
