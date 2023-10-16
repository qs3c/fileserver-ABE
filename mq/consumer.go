package mq

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

var done chan bool

//var msgs chan

var msgs <-chan amqp.Delivery = nil

// StartConsume : 接收消息
func StartConsume(qName, cName string, callback func(msg []byte) bool) {
	if msgs == nil {
		log.Println("第一次生成 msgs")
		msgs, _ = channel.Consume(
			qName,
			cName,
			true,  //自动应答
			false, // 非唯一的消费者
			false, // rabbitMQ只能设置为false
			false, // noWait, false表示会阻塞直到有消息过来
			nil)
	} else {
		log.Println("非第一次沿用以前的msgs", msgs)
	}

	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	//for a := range msgs {
	//	log.Println("拿到 msg了！", string(a.Body))
	//} //msg 拿一次 publish 就拿空了，就阻塞住了！所以后面的没执行
	//log.Println(callback)

	done = make(chan bool)

	//问题好像处在这里用 goroutine 开协程异步获取数据，开启以后不会停止啊
	//所以后面的第二次就取不到数据了
	//都被第一次的抢走了，第一次开启的goroutin还在不断的消费py传来的信息
	//想要一发一收不应该用这种异步的持续消费
	//好吧问题不在于这里，因为新开go协程，这个函数推出了协程就终止了，不存在函数推出了协程还在持续消费的情况
	//所以问题应该是出在

	log.Println("从复用的msgs中读取数据")
	go func() {
		// 循环读取channel的数据
		log.Println("协程开始了！")
		for d := range msgs {
			processErr := callback(d.Body)
			time.Sleep(time.Second)
			//log.Println("执行到这了吗，我想请问一下？")
			if !processErr {
				// TODO: 将任务写入错误队列，待后续处理
				log.Println("回调函数Transfer出现错误！")
			}
		}
		log.Println("协程退出了！")
	}()

	// 接收done的信号, 没有信息过来则会一直阻塞，避免该函数退出
	<-done

	// 关闭通道
	// 问题就出在这一句了！！！ close nil channel ！
	//channel.Close()
}

// StopConsume : 停止监听队列
func StopConsume() {
	done <- true
}

func StartConsume_nooldmsgs(qName, cName string, callback func(msg []byte) bool) {

	msgs, _ = channel.Consume(
		qName,
		cName,
		true,  //自动应答
		false, // 非唯一的消费者
		false, // rabbitMQ只能设置为false
		false, // noWait, false表示会阻塞直到有消息过来
		nil)

	done = make(chan bool)

	go func() {
		// 循环读取channel的数据
		for d := range msgs {
			processErr := callback(d.Body)
			//log.Println("执行到这了吗，我想请问一下？")
			if !processErr {
				// TODO: 将任务写入错误队列，待后续处理
				log.Println("回调函数Transfer出现错误！")
			}
		}
	}()

	// 接收done的信号, 没有信息过来则会一直阻塞，避免该函数退出
	<-done

}

func StartConsumeOne(qName, cName string, callback func(msg []byte) bool) {
	if msgs == nil {
		log.Println("第一次生成 msgs")
		msgs, _ = channel.Consume(
			qName,
			cName,
			true,  //自动应答
			false, // 非唯一的消费者
			false, // rabbitMQ只能设置为false
			false, // noWait, false表示会阻塞直到有消息过来
			nil)
		//log.Println("第一次，所以启动从msgs获取数据的协程，这个协程会持续运行，更新全局变量Key。")
	} else {
		log.Println("非第一次沿用以前的msgs", msgs)
		log.Println("什么都不需要做捏！")
	}
	data := <-msgs
	processErr := callback(data.Body)
	time.Sleep(time.Second)
	//log.Println("执行到这了吗，我想请问一下？")
	if !processErr {
		// TODO: 将任务写入错误队列，待后续处理
		log.Println("回调函数Transfer出现错误！")
	}

}
