package initialize

import (
	"finders-server/global"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func MySql() {
	admin := global.CONFIG.MySQL
	connect_str := admin.Username + ":" + admin.Password + "@(" + admin.Addr + ")/" + admin.Database + "?" + admin.Config
	if db, err := gorm.Open("mysql", connect_str); err != nil {
		global.LOG.Fatalf("MySQl连接异常 ERROR: %s", err)
	} else {
		db.DB().SetMaxIdleConns(admin.MaxIdleConns)
		db.DB().SetMaxOpenConns(admin.MaxOpenConns)
		global.DB = db
	}

}
