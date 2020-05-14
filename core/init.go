/*
core 最先加载config
*/
package core

import (
	"finders-server/config"
	"finders-server/global"
)

func init() {

	//初始化读入配置文件
	global.CONFIG = config.InitConfig()

	//初始化日志
	global.LOG = InitLogger(global.CONFIG.Log)

	global.LOG.Debug("Core init")

}
