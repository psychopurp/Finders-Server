package initialize

import (
	"finders-server/global"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func MySql() {
	admin := global.CONFIG.MySQL
	connectStr := admin.Username + ":" + admin.Password + "@(" + admin.Addr + ")/" + admin.Database + "?" + admin.Config
	if db, err := gorm.Open("mysql", connectStr); err != nil {
		global.LOG.Fatalf("MySQl连接异常 ERROR: %s", err)
	} else {
		db.DB().SetMaxIdleConns(admin.MaxIdleConns)
		db.DB().SetMaxOpenConns(admin.MaxOpenConns)
		// db.LogMode(global.CONFIG.MySQL.LogMode)
		global.DB = db
	}

}
