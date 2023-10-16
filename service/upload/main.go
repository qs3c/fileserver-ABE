package main

import (
	"fileserver_enc/common"
	"fileserver_enc/mq"
	dbproxy "fileserver_enc/service/dbproxy/client"
	cfg "fileserver_enc/service/upload/config"
	upProto "fileserver_enc/service/upload/proto"
	"fileserver_enc/service/upload/route"
	upRpc "fileserver_enc/service/upload/rpc"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"
	"time"
)

func startRPCService() {
	myRegistry := consul.NewRegistry(registry.Addrs("172.27.103.231:8500"))
	service := micro.NewService(
		micro.Name("go.micro.service.upload"), // 服务名称
		micro.RegisterTTL(time.Second*10),     // TTL指定从上一次心跳间隔起，超过这个时间服务会被服务发现移除
		micro.RegisterInterval(time.Second*5), // 让服务在指定时间内重新注册，保持TTL获取的注册时间有效
		micro.Flags(common.CustomFlags...),
		micro.Registry(myRegistry),
	)
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 检查是否指定mqhost
			mqhost := c.String("mqhost")
			if len(mqhost) > 0 {
				log.Println("custom mq address: " + mqhost)
				mq.UpdateRabbitHost(mqhost)
			}
			return nil
		}),
	)

	// 初始化dbproxy client
	dbproxy.Init(service)
	// 初始化mq client
	mq.Init()

	upProto.RegisterUploadServiceHandler(service.Server(), new(upRpc.Upload))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startAPIService() {
	router := route.Router()
	router.Run(cfg.UploadServiceHost)
}

func main() {
	//os.MkdirAll(config.TempLocalRootDir, 0777)
	//os.MkdirAll(config.TempPartRootDir, 0777)

	// api 服务
	go startAPIService()

	// rpc 服务
	startRPCService()
}
