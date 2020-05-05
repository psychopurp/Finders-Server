package main

import (
	"finders-server/core"
	"finders-server/global"
	"finders-server/initialize"
	"fmt"
)

func main() {

	switch global.CONFIG.System.DB {
	case "mysql":
		global.LOG.Debug("in mysql")
	default:
		fmt.Println("default")
	}
	initialize.Routers()

	core.RunServer()
}
