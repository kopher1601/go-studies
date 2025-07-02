package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/kopher1601/go-studies/go-gateway/config"
)

const (
	_allAcks = "all"
)

type Producer struct {
	cfg      config.Producer
	producer *kafka.Producer
}

func NewProducer(cfg config.Producer) Producer {
	url := cfg.URL
	id := cfg.ClientID
	acks := cfg.Acks

	if acks == "" {
		acks = _allAcks
	}

	conf := &kafka.ConfigMap{
		"bootstrap.servers": url, // kafka client url
		"client.id":         id,
		"acks":              acks,
	}

	producer, err := kafka.NewProducer(conf)
	if err != nil {
		panic(err)
	}

	return Producer{
		cfg:      cfg,
		producer: producer,
	}
}

func (p Producer) SendEvent(v []byte) {
	topic := p.cfg.Topic

	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: v,
	}, nil)

	if err != nil {
		log.Println("failed to send event to kafka", string(v))
	} else {
		log.Println("event sent to kafka", string(v))
	}
}
