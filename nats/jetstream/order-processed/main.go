package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	log.Println("开始处理新订单 ...")
	sub, _ := js.PullSubscribe("ORDERS.received", "NEW")

	for {
		msgs, _ := sub.Fetch(100)
		for _, msg := range msgs {
			log.Println("正在处理新订单： " + string(msg.Data))
			// TODO：
			_, err := js.Publish("ORDERS.processed", []byte(string(msg.Data)+",processed"))
			if err != nil {
				log.Println(err.Error())
				continue
			}
			msg.Ack()
			log.Println("新订单处理完毕： " + string(msg.Data))
		}
	}
}
