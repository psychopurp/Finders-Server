package st

import (
	"finders-server/global"
	"runtime"
)

func Debug(arg ...interface{}) {
	global.LOG.Debug(arg...)
}

func printCallerFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func DebugWithFuncName(err error) {
	pc, file, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	errInfo := "check error"
	if err != nil {
		errInfo = err.Error()
	}
	Debug("funcName:", funcName, "file:", file, "line:", line, " err: ", errInfo)

}
