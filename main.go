package main

import (
	"finders-server/core"
	"finders-server/global"
	"finders-server/initialize"
	"fmt"
)

var (
	_ = core.InitLogger
	_ = global.LOG
	_ = initialize.MySql
)

func test() {
	fmt.Println(global.CONFIG)

}

func main() {
	// test()
	global.LOG.Debug("初始化redis")
	err := initialize.Redis()
	if err != nil {
		global.LOG.Debug("redis初始化失败")
	}
	switch global.CONFIG.System.DB {
	case "mysql":
		initialize.MySql()
	case "sqlite":
		initialize.Sqlite()
	default:
		fmt.Println("default")
	}
	global.LOG.Debug("连接数据库")
	initialize.DBTables()

	//程序结束前关闭数据库连接
	defer global.DB.Close()

	core.RunServer()
}
