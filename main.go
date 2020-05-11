package main

import (
	"finders-server/core"
	"finders-server/global"
	"finders-server/initialize"
	"finders-server/model"
	"fmt"
)

var (
	_ = core.InitConfig
	_ = global.LOG
	_ = initialize.MySql
)

func test() {
	db := global.DB
	u := model.User{}
	// u.Phone = "52"
	// u.UserName = "obott90"
	err := db.Model(&u).First(&u).Error
	// u, err := service.Register(u)
	if err != nil {
		fmt.Println("check :", err)
	}
	fmt.Println(u.UserInfo, u.Relations)

}

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

	// test()
	core.RunServer()
}
