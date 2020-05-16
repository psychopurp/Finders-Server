package initialize

// package main

import (
	"finders-server/global"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Sqlite() {

	admin := global.CONFIG.SQLite
	if db, err := gorm.Open("sqlite3", fmt.Sprintf("%s?%s", admin.Path, admin.Config)); err != nil {
		global.LOG.Fatalf("Sqlite连接异常 ERROR: %s", err)
	} else {
		global.DB = db
		global.DB.LogMode(admin.LogMode)
	}
}
