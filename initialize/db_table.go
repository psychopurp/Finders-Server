package initialize

import (
	"finders-server/global"
	"finders-server/model"
)

//注册数据库表专用
func DBTables() {
	db := global.DB
	db.AutoMigrate(&model.User{}, &model.UserInfo{}, &model.Relation{})
	db.AutoMigrate(&model.Login{})
	db.AutoMigrate(&model.Admin{})
	db.AutoMigrate(&model.Media{})
	global.LOG.Debug("Register tables success")
}
