/*
项目核心---日志模块
需要实现的功能：

1.以直观简洁的方式输出日志
2.能够将日志保存到文件里
3.能够将日志保存到数据库里
4.能够日后对日志进行分析

logrus六种日志级别：debug, info, warn, error, fatal, panic

*/

package core

import (
	"finders-server/config"
	"fmt"
	"io"
	"os"
	"strings"

	oplogging "github.com/op/go-logging"
)

const (
	logDir      = "log"
	logSoftLink = "latest_log"
	module      = "finders-server"
)

var (
	defaultFormatter = `%{color:bold} %{time:2006/01/02 15:04:05} %{longfile} ▶ [%{level:.6s}] %{message}%{color:reset}`
)

func InitLogger(config config.LogConfig) *oplogging.Logger {
	c := config
	if c.Prefix == "" {
		_ = fmt.Errorf("Logger prefix not fount")
	}
	logger := oplogging.MustGetLogger(module)
	var backends []oplogging.Backend

	backends = registerStdout(c, backends)
	oplogging.SetBackend(backends...)
	return logger

}

func registerStdout(c config.LogConfig, backends []oplogging.Backend) []oplogging.Backend {
	if c.Stdout != "" {
		level, err := oplogging.LogLevel(c.Stdout)
		if err != nil {
			fmt.Println(err)
		}
		backends = append(backends, createBackend(os.Stdout, c, level))
	}
	return backends
}

func createBackend(out io.Writer, c config.LogConfig, level oplogging.Level) oplogging.Backend {
	backend := oplogging.NewLogBackend(out, c.Prefix, 0)
	stdOutWriter := false
	if out == os.Stdout {
		stdOutWriter = true
	}
	format := getLogFormatter(c, stdOutWriter)
	backendLeveled := oplogging.AddModuleLevel(oplogging.NewBackendFormatter(backend, format))
	backendLeveled.SetLevel(level, module)
	return backendLeveled

}

func getLogFormatter(c config.LogConfig, stdOutWriter bool) oplogging.Formatter {
	pattern := defaultFormatter
	if !stdOutWriter {
		//color is only required for console output
		pattern = strings.Replace(pattern, "%{color:bold}", "", -1)
		pattern = strings.Replace(pattern, "%{color:reset}", "", -1)
	}
	if !c.Logfile {
		//remove %{logfile} tag
		pattern = strings.Replace(pattern, "%{longfile}", "", -1)
	}
	return oplogging.MustStringFormatter(pattern)

}
