package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/yanuar-nc/go-boiler-plate/src/shared/domain"
	"github.com/yanuar-nc/golang/helper"
)

// KafkaMessageBroker struct
type Kafka struct {
	config  *sarama.Config
	brokers []string
}

func NewKafka(brokers []string, config *sarama.Config) *Kafka {
	// config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	// config.ClientID = clientId
	return &Kafka{
		config:  config,
		brokers: brokers,
	}
}

func (k *Kafka) Publish(ctx context.Context, topic string, event domain.EventMessage) error {
	producer, err := sarama.NewSyncProducer(k.brokers, k.config)
	if err != nil {
		return err
	}

	msgJson, _ := json.Marshal(event)

	message := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(msgJson),
	}
	message.Key = sarama.StringEncoder(event.Key)
	message.Timestamp = time.Now()

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		return err
	}

	helper.Log(log.InfoLevel, fmt.Sprintf("Producer is successfully in partition %d and offset %d", partition, offset), "Kafka", "producer_send_message")
	return nil
}
