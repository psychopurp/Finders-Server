/*
core 最先加载config
*/
package core

import (
	"finders-server/config"
	"finders-server/global"
	"flag"
)

var (
	config_root = flag.String("conf", "conf", "Input config dir")
	log_dir     = flag.String("log", "", "Input log dir")
)

func init() {
	flag.Parse()
	//初始化读入配置文件
	global.CONFIG = config.InitConfig(*config_root)
	//初始化日志
	global.LOG = InitLogger(global.CONFIG.Log)

	global.LOG.Debug("Core init")

}
