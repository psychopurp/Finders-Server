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
	"path"
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
	// 检查是否有前缀
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
		backends = append(backends, getFileBackend(c))
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




func  getFileBackend(c config.LogConfig) oplogging.LeveledBackend {
	//判断是否存在该文件夹
	if err := os.MkdirAll(logDir, 0777); err != nil {
		panic(err)
	}
	// 打开一个文件
	file, err := os.OpenFile(path.Join(logDir, module+"_info.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	//backend := l.getLogBackend(file, LogLevelMap[l.level])
	level, err := oplogging.LogLevel(c.Stdout)
	if err != nil {
		panic(err)
	}
	backend := getLogBackend(c, file, int(level))
	//logging.SetBackend(backend)
	return backend
}


func  getLogBackend(c config.LogConfig, out io.Writer, level int) oplogging.LeveledBackend {
	pattern := defaultFormatter
	pattern = strings.Replace(pattern, "%{color:bold}", "", -1)
	pattern = strings.Replace(pattern, "%{color:reset}", "", -1)
	if !c.Logfile {
		//remove %{logfile} tag
		pattern = strings.Replace(pattern, "%{longfile}", "", -1)
	}
	backend := oplogging.NewLogBackend(out, c.Prefix, 1)
	format := oplogging.MustStringFormatter(pattern)
	backendFormatter := oplogging.NewBackendFormatter(backend, format)
	backendLeveled := oplogging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(oplogging.Level(level), "")
	return backendLeveled
}
