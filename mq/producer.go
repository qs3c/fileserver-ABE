package mq

import (
	"fileserver_enc/config"
	"log"

	"github.com/streadway/amqp"
)

// Publish : 发布消息
func Publish(exchange, routingKey string, msg []byte) bool {

	if !initChannel(config.RabbitURL) {
		log.Println("初始化channel失败！")
		return false
	}

	if nil == channel.Publish(
		exchange,
		routingKey,
		false, // 如果没有对应的queue, 就会丢弃这条消息
		false, //
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg}) {
		return true
	} else {
		log.Println("channel.Publish执行失败！")
		return false
	}

}
