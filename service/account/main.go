package main

import (
	"fileserver_enc/common"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"log"
	"time"
	// k8s "github.com/micro/kubernetes/go/micro"
	"go-micro.dev/v4/registry"
	//"fileserver_enc/common"
	"fileserver_enc/service/account/handler"
	proto "fileserver_enc/service/account/proto"
	dbproxy "fileserver_enc/service/dbproxy/client"
)

func main() {
	myRegistry := consul.NewRegistry(registry.Addrs("172.27.103.231:8500"))
	service := micro.NewService(
		// service := k8s.NewService(
		micro.Name("go.micro.service.user"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
		micro.Registry(myRegistry),
	)

	// 初始化service, 解析命令行参数等
	service.Init()

	// 初始化dbproxy client
	dbproxy.Init(service)

	proto.RegisterUserServiceHandler(service.Server(), new(handler.User))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
