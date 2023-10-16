package test

import (
	"bytes"
	"fileserver_enc/mq"
	"fileserver_enc/service/abe/config"
	"fileserver_enc/service/abe/handler"
	"log"
	"os/exec"
)

func AbEncryption(accessPolicy string) {
	//1.拿到访问结构
	//accessPolicy := in.AccessPolicy
	//pubData,_ := json.Marshal(accessPolicy)
	pubData := []byte(accessPolicy)
	//2.发送给 py
	ifSuc := mq.Publish("python.go.trans", "gotopy", pubData)
	if !ifSuc {
		// TODO: 当前发送转移信息失败，稍后重试
		log.Println("当前发送转移信息失败，稍后重试")
	} else {
		log.Println("成功publish等待消费！")
	}
	//time.Sleep(time.Second)
	// 执行py脚本（先手动）
	cmd := exec.Command("python3", "/root/ABE-master/test/aes_key_gen.py")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println("执行py脚本出错！")
		log.Println(err.Error(), stderr.String())

	}
	log.Println("python执行结果:",out)
	//time.Sleep(time.Second)
	// 从mq获取结果
	mq.Init()
	log.Println("开始消费...")
	mq.StartConsumeOne(
		"python.go.trans.pytogo",
		"transfer_key",
		handler.HandleMessage) // 拿到 message后的处理逻辑在 HandleMessage这里
	//拿到 key 就停止
	//mq.StopConsume() ,应该在HandleMessage中 stop
	log.Println("此时的key为：", config.Key)

}
