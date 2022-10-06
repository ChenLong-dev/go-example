package rabbitmq

import (
	"database/rabbitmq/consumer"
	"database/rabbitmq/producer"
)

func RabbitmqClient() {
	go func() {
		consumer.Consumer()
	}()

	go func() {
		producer.Producter()
	}()

	select {}
}
