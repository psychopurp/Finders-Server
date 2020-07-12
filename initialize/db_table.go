package initialize

import (
	"finders-server/global"
	"finders-server/model"
)

//注册数据库表专用
func DBTables() {
	db := global.DB
	db.AutoMigrate(&model.User{}, &model.UserInfo{}, &model.Relation{})
<<<<<<< HEAD
	// db.AutoMigrate(&model.UserTest{})
=======
	db.AutoMigrate(&model.Login{})
>>>>>>> test
	global.LOG.Debug("Register tables success")
}
