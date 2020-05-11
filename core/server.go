package core

import (
	"finders-server/global"
	"finders-server/initialize"
	"fmt"
	"net/http"
	"time"
)

func RunServer() {

	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Port)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	global.LOG.Debugf("%s Serve run success on http://%s:%d ", global.CONFIG.AppName, global.CONFIG.System.IP, global.CONFIG.System.Port)

	global.LOG.Error(s.ListenAndServe())

}
