package mq

import (
	"fileserver_enc/config"
	"log"
	"fmt"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel

// 如果异常关闭，会接收通知
var notifyClose chan *amqp.Error

// UpdateRabbitHost : 更新mq host
func UpdateRabbitHost(host string) {
	//直接运行用这个：
	config.RabbitURL = fmt.Sprintf("amqp://guest:guest@%s:5672/", host)
	//docker有启动参数用这个：
	//config.RabbitURL = host
}

// Init : 初始化MQ连接信息
func Init() {
	// 是否开启异步转移功能，开启时才初始化rabbitMQ连接
	if !config.AsyncTransferEnable {
		return
	}
	if initChannel(config.RabbitURL) {
		channel.NotifyClose(notifyClose)
	}
	// 断线自动重连
	go func() {
		for {
			select {
			case msg := <-notifyClose:
				conn = nil
				channel = nil
				log.Printf("onNotifyChannelClosed: %+v\n", msg)
				initChannel(config.RabbitURL)
			}
		}
	}()
}

func initChannel(rabbitHost string) bool {
	if channel != nil {
		return true
	}
	log.Println(rabbitHost)
	conn, err := amqp.Dial(rabbitHost)
	if err != nil {
		log.Println(err.Error())
		log.Println("Dial阶段出问题！")
		return false
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Println(err.Error())
		log.Println("产生channel出问题！")
		return false
	}
	return true
}
