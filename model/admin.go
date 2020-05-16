package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

/*
DB Table Details


CREATE TABLE `admins` (
  `admin_id` varchar(30) NOT NULL COMMENT '管理员ID',
  `admin_name` varchar(30) NOT NULL COMMENT '管理员名称',
  `admin_password` varchar(30) NOT NULL COMMENT '管理员密码',
  `permission` int(11) NOT NULL COMMENT '管理员权限',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '注册时间',
  `deleted_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '注销管理员的时间',
  PRIMARY KEY (`admin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Admin struct is a row record of the admins table in the employees database
type Admin struct {
	AdminID       string    `gorm:"column:admin_id;type:VARCHAR;size:30;primary_key" json:"admin_id"`  //[ 0] admin_id                                       VARCHAR[30]          null: false  primary: true   auto: false
	AdminName     string    `gorm:"column:admin_name;type:VARCHAR;size:30;" json:"admin_name"`         //[ 1] admin_name                                     VARCHAR[30]          null: false  primary: false  auto: false
	AdminPassword string    `gorm:"column:admin_password;type:VARCHAR;size:30;" json:"admin_password"` //[ 2] admin_password                                 VARCHAR[30]          null: false  primary: false  auto: false
	Permission    int       `gorm:"column:permission;type:INT;" json:"permission"`                     //[ 3] permission                                     INT                  null: false  primary: false  auto: false
	CreatedAt     time.Time `gorm:"column:created_at;type:DATETIME;" json:"created_at"`                //[ 4] created_at                                     DATETIME             null: false  primary: false  auto: false
	DeletedAt     null.Time `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`                //[ 5] deleted_at                                     DATETIME             null: true   primary: false  auto: false

}

// TableName sets the insert table name for this struct type
func (a *Admin) TableName() string {
	return "admins"
}

func (a *Admin) BeforeSave() error {
	return nil
}

func (a *Admin) Prepare() {
}

func (a *Admin) Validate(action Action) error {

	return nil
}
