package initialize

import (
	"finders-server/global"
<<<<<<< HEAD
=======
	"time"
>>>>>>> test

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
<<<<<<< HEAD
	}

}
=======

		// 取消注释后在更新和创建时会更新时间  更新时会更新时间
		//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
		//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	}

}


func updateTimeStampForCreateCallback(scope *gorm.Scope){
	if !scope.HasError(){
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok{
			if createTimeField.IsBlank{
				createTimeField.Set(nowTime)
			}
		}

		if updateTimeField, ok := scope.FieldByName("UpdatedAt"); ok{
			if updateTimeField.IsBlank{
				updateTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope){
	if _, ok := scope.Get("gorm:update_column"); !ok{
		scope.SetColumn("UpdatedAt", time.Now().Unix())
	}
}
>>>>>>> test
