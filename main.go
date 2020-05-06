package main

import (
	"finders-server/core"
	"finders-server/global"
	"finders-server/initialize"
	"fmt"
)

func main() {

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
