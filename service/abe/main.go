package main

import (
	"fileserver_enc/common"
	"fileserver_enc/service/abe/handler"
	proto "fileserver_enc/service/abe/proto"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"
	"time"
)

func main() {
	//测试
	//access := "((ONE and THREE) and (TWO OR FOUR))#/root/ABE-master/test/temp"
	//test.AbEncryption(access)

	//做成服务
	myRegistry := consul.NewRegistry(registry.Addrs("localhost:8500"))
	service := micro.NewService(
		micro.Name("go.micro.service.abe"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
		micro.Registry(myRegistry),
	)

	// 初始化service, 解析命令行参数等
	service.Init()

	proto.RegisterEncUploadServiceHandler(service.Server(), new(handler.Abe))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
