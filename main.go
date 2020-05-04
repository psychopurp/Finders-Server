package main

import (
	"finders-server/core"
	"finders-server/global"
	"fmt"
)

func main() {
	test(5, 6, "elyar")
	// model.Test2()
	fmt.Println(global.CONFIG.Log)

	switch global.CONFIG.System.DB {
	case "mysql":
		fmt.Println("this is in case")
	default:
		fmt.Println("default")
	}

	core.RunServer()
}
