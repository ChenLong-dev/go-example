package main

import (
	"database/kafka"
)

func main() {
	//etcd.EtcdClient()
	//redis.RedisClient()
	//mysql.MysqlClient()
	//rabbitmq.RabbitmqClient()
	kafka.KafkaClient()
}
