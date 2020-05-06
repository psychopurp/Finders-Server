/*
core 最先加载config
*/
package core

import "finders-server/global"

func init() {

	InitConfig()
	InitLogger()
	global.LOG.Debug("Core init")

}
