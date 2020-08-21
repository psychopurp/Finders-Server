package st

import "finders-server/global"

func Debug(args ...interface{}) {
	global.LOG.Debug(args...)
}
