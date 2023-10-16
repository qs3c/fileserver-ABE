package process

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"fileserver_enc/mq"
	dbcli "fileserver_enc/service/dbproxy/client"
	"fileserver_enc/store/oss"
)

// Transfer : 处理文件转移
// callback函数定义：拿到数据byte流后怎么处理
func Transfer(msg []byte) bool {
	log.Println(string(msg))

	//1.创好消息体，消息byte流去json化放入结构体
	pubData := mq.TransferData{}
	err := json.Unmarshal(msg, &pubData)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	//处理消息体中数据
	fin, err := os.Open(pubData.CurLocation)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	err = oss.Bucket().PutObject(
		pubData.DestLocation,
		bufio.NewReader(fin))
	if err != nil {
		log.Println(err.Error())
		return false
	}

	resp, err := dbcli.UpdateFileLocation(
		pubData.FileHash,
		pubData.DestLocation)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if !resp.Suc {
		log.Println("更新数据库异常，请检查:" + pubData.FileHash)
		return false
	}
	return true
}
