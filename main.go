package main

import (
	"finders-server/core"
	"finders-server/global"
	"finders-server/initialize"
	"finders-server/model"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

var (
	_ = core.InitConfig
	_ = global.LOG
	_ = initialize.MySql
)

func test() {
	db := global.DB
	uid := uuid.NewV4()
	u := new(model.User)
	u.UserID = uid
	u.Status = model.Normal
	u.Phone = "5200"
	u2 := new(model.User)
	fmt.Println(u2)
	db.First(u2)
	fmt.Println(u2.UserInfo)
	fmt.Println(u.UserID, len(u.UserID))
	err := db.Create(&u).Error
	fmt.Println(err)

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

	test()
	// core.RunServer()
}
