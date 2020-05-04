/*
项目日志模块
需要实现的功能：

1.以直观简洁的方式输出日志
2.能够将日志保存到文件里
3.能够将日志保存到数据库里
4.能够日后对日志进行分析

*/

package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	logDir      = "log"
	logSoftLink = "latest_log"
	module      = "finders-server"
)

var (
	defaultFormatter = `%{time:2006/01/02 - 15:04:05.000} %{longfile} %{color:bold}▶ [%{level:.6s}] %{message}%{color:reset}`
)

func main() {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.Warning("this is test")
}
