package main

import (
	"fileserver_enc/common"
	"fileserver_enc/service/dbproxy/config"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	//"go-micro.dev/v4/config/source/cli"
	//"github.com/go-micro/cli"
	"github.com/urfave/cli/v2"
	"log"
	"time"

	dbConn "fileserver_enc/service/dbproxy/conn"
	dbProxy "fileserver_enc/service/dbproxy/proto"
	dbRpc "fileserver_enc/service/dbproxy/rpc"
)

func startRpcService() {
	myRegistry := consul.NewRegistry(registry.Addrs("172.27.103.231:8500"))
	service := micro.NewService(
		micro.Name("go.micro.service.dbproxy"), // 在注册中心中的服务名称
		micro.RegisterTTL(time.Second*10),      // 声明超时时间, 避免consul不主动删掉已失去心跳的服务节点
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
		micro.Registry(myRegistry),
	)

	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 检查是否指定dbhost
			dbhost := c.String("dbhost")
			if len(dbhost) > 0 {
				log.Println("custom db address is : " + dbhost)
				config.UpdateDBHost(dbhost)
			}
			return nil
		}),
	)

	// 初始化db connection
	dbConn.InitDBConn()

	dbProxy.RegisterDBProxyServiceHandler(service.Server(), new(dbRpc.DBProxy))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func main() {
	startRpcService()
}

// res, err := mapper.FuncCall("/user/UserExist", []interface{}{"haha"}...)
// log.Printf("error: %+v\n", err)
// log.Printf("result: %+v\n", res[0].Interface())

// res, err = mapper.FuncCall("/user/UserExist", []interface{}{"admin"}...)
// log.Printf("error: %+v\n", err)
// log.Printf("result: %+v\n", res[0].Interface())
