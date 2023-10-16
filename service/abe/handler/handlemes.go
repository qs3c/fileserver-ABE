package handler

import (
	//"fileserver_enc/mq"
	"fileserver_enc/service/abe/config"
	"log"
)

func HandleMessage(msg []byte) bool {
	//对数据没什么处理，就是放入全局变量key中
	//log.Println(string(msg))
	log.Println("修改了key")
	config.Key = string(msg)
	//mq.StopConsume()
	return true
}
