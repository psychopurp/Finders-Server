package initialize

import "finders-server/global"

//注册数据库表专用
func DBTables() {
	db := global.DB
	db.AutoMigrate()
	global.LOG.Debug("Register tables success")
}
