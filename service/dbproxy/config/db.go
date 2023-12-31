package config

import "fmt"

var (
	// MySQLSource : 要连接的数据库源；
	// 其中test:test 是用户名密码；
	// 127.0.0.1:3306 是ip及端口；
	// fileserver 是数据库名;
	// charset=utf8 指定了数据以utf8字符编码进行传输
	MySQLSource = "root:root@tcp(172.27.103.231:3339)/fileserver?charset=utf8"
)

func UpdateDBHost(host string) {
	MySQLSource = fmt.Sprintf("root:root@tcp(%s)/fileserver?charset=utf8", host)
}
