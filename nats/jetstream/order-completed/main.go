package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"strings"
)

func main() {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	log.Println("开始处理订单 ...")
	sub, _ := js.PullSubscribe("ORDERS.processed", "DISPATCH")
	for {
		msgs, _ := sub.Fetch(10)
		for _, msg := range msgs {
			orderInfo := strings.Split(string(msg.Data), ",")
			orderNo := orderInfo[0]
			log.Println("处理订单： " + orderNo)
			// TODO：
			_, err := js.Publish("ORDERS.completed", []byte(orderNo+",completed"))
			if err != nil {
				log.Println(err.Error())
				continue
			}
			log.Println("订单处理完毕： " + orderNo)
			msg.Ack()
		}
	}
}
