package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"time"

	"go-micro.dev/v4"
	//_ "github.com/micro/go-plugins/registry/kubernetes"

	"fileserver_enc/common"
	dbproxy "fileserver_enc/service/dbproxy/client"
	cfg "fileserver_enc/service/download/config"
	dlProto "fileserver_enc/service/download/proto"
	"fileserver_enc/service/download/route"
	dlRpc "fileserver_enc/service/download/rpc"
)

func startRPCService() {
	myRegistry := consul.NewRegistry(registry.Addrs("172.27.103.231:8500"))
	service := micro.NewService(
		micro.Name("go.micro.service.download"), // 在注册中心中的服务名称
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
		micro.Registry(myRegistry),
	)
	service.Init()

	// 初始化dbproxy client
	dbproxy.Init(service)

	dlProto.RegisterDownloadServiceHandler(service.Server(), new(dlRpc.Download))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startAPIService() {
	router := route.Router()
	router.Run(cfg.DownloadServiceHost)
}

func main() {
	// api 服务
	go startAPIService()

	// rpc 服务
	startRPCService()
}
