/*
项目核心---日志模块
需要实现的功能：

1.以直观简洁的方式输出日志
2.能够将日志保存到文件里
3.能够将日志保存到数据库里
4.能够日后对日志进行分析

logrus六种日志级别：debug, info, warn, error, fatal, panic

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
	file, err := os.OpenFile("logrus.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		logger.Info("Faild to open log file")
	}
	logger.Out = file
	logger.Warning("this is test")
}
