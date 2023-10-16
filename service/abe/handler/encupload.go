package handler

import (
	"bytes"
	"fileserver_enc/mq"
	"fileserver_enc/service/abe/config"
	proto "fileserver_enc/service/abe/proto"
	"log"
	"os/exec"
)
import "context"

type Abe struct{}

func (a *Abe) AbEncryption(ctx context.Context, in *proto.ReqAccessPolicy, out *proto.RespKey) error {
	//1.拿到访问结构(#后带上了地址)
	log.Println("开始执行rpc服务的函数！")
	accessPolicy := in.AccessPolicy
	//pubData,_ := json.Marshal(accessPolicy)
	pubData := []byte(accessPolicy)
	//2.发送给 py
	ifSuc := mq.Publish("python.go.trans", "gotopy", pubData)
	if !ifSuc {
		log.Println("当前发送转移信息失败，稍后重试")
		return nil
	} else {
		log.Println("成功publish等待消费！")
	}
	// 执行py脚本（先手动）
	cmd := exec.Command("python3", "/root/ABE-master/test/aes_key_gen.py")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println("执行py脚本出错！")
		log.Println(err.Error(), stderr.String())

	}
	// 从mq获取结果
	mq.Init()
	log.Println("开始消费...")
	mq.StartConsume(
		"python.go.trans.pytogo",
		"transfer_key",
		HandleMessage) // 拿到 message后的处理逻辑在 HandleMessage这里
	//拿到 key 就停止
	//mq.StopConsume() ,应该在HandleMessage中 stop
	out.Key = config.Key
	return nil
}
