package kafka

import (
	"database/kafka/consumer"
	"database/kafka/producer"
)

func KafkaClient() {
	go func() {
		consumer.Consumer()
	}()

	go func() {
		producer.Porducer()
	}()

	select {}
}
