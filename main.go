package main

import (
	"log"

	"github.com/google/uuid"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	publisher, err := kafka.NewPublisher(kafka.PublisherConfig{
		Brokers:               []string{"localhost:9092", "localhost:9091"},
		Marshaler:             kafka.DefaultMarshaler{},
		OverwriteSaramaConfig: kafka.DefaultSaramaSyncPublisherConfig(),
		OTELEnabled:           false,
	},
		watermill.NewStdLogger(true, true),
	)

	if err != nil {
		log.Fatal(err)
	}

	err = publisher.Publish("testtopic", message.NewMessage(uuid.NewString(), []byte("hello")))

	if err != nil {
		log.Fatal(err)
	}
}
