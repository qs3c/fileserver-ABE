package common

//import "github.com/go-micro/cli"
import "github.com/urfave/cli/v2"

// CustomFlags : 自定义命令行参数
var CustomFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "dbhost",
		Value: "localhost",
		Usage: "database address",
	},
	&cli.StringFlag{
		Name:  "mqhost",
		Value: "localhost",
		Usage: "mq(rabbitmq) address",
	},
	&cli.StringFlag{
		Name:  "cachehost",
		Value: "localhost",
		Usage: "cache(redis) address",
	},
	&cli.StringFlag{
		Name:  "cephhost",
		Value: "localhost",
		Usage: "ceph address",
	},
}
