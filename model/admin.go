package model

import (
	"database/sql"
	"finders-server/global"
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
	AdminID       string     `gorm:"column:admin_id;type:varchar(50);primary_key" json:"admin_id"`                           //[ 0] admin_id                                       VARCHAR[30]          null: false  primary: true   auto: false
	AdminName     string     `gorm:"column:admin_name;type:varchar(30);" json:"admin_name"`                                  //[ 1] admin_name                                     VARCHAR[30]          null: false  primary: false  auto: false
	AdminPassword string     `gorm:"column:admin_password;type:varchar(30);" json:"admin_password"`                          //[ 2] admin_password                                 VARCHAR[30]          null: false  primary: false  auto: false
	AdminPhone    string     `gorm:"column:admin_phone;type:varchar(30);unique_index:unique_admin_phone" json:"admin_phone"` //[ 2] admin_password                                 VARCHAR[30]          null: false  primary: false  auto: false
	Permission    int        `gorm:"column:permission;type:INT;" json:"permission"`                                          //[ 3] permission                                     INT                  null: false  primary: false  auto: false
	CreatedAt     time.Time  `gorm:"column:created_at;type:DATETIME;" json:"created_at"`                                     //[ 4] created_at                                     DATETIME             null: false  primary: false  auto: false
	UpdatedAt     time.Time  `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"`                                     //[16] updated_at                                     DATETIME             strue   primary: false  auto: false
	DeletedAt     *time.Time `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`                                     //[ 5] deleted_at                                     DATETIME             null: true   primary: false  auto: false
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

const (
	SUPER  = 1
	NORMAL = 2
)

func GetAdminByAdminName(adminName string) (admin Admin, err error) {
	db := global.DB
	err = db.Where("admin_name = ?", adminName).First(&admin).Error
	return
}

func GetAdminByAdminID(adminID string) (admin Admin, err error) {
	db := global.DB
	err = db.Where("admin_id = ?", adminID).First(&admin).Error
	return
}
func ExistAdminByAdminNameAndPassword(name, password string) (admin Admin, isExist bool) {
	db := global.DB
	data := make(map[string]interface{})
	data["admin_name"] = name
	data["admin_password"] = password
	isExist = !db.Where(data).First(&admin).RecordNotFound()
	return
}

func ExistAdminByPhone(phone string) (admin Admin, isExist bool) {
	db := global.DB
	data := make(map[string]interface{})
	data["admin_phone"] = phone
	isExist = !db.Where(data).First(&admin).RecordNotFound()
	return
}

func AddAdmin(admin *Admin) (err error) {
	db := global.DB
	err = db.Create(admin).Error
	return
}

func AddAdminByMaps(data map[string]interface{}) (admin Admin, err error) {
	db := global.DB
	admin = Admin{
		AdminID:       data["admin_id"].(string),
		AdminName:     data["admin_name"].(string),
		AdminPassword: data["admin_password"].(string),
		AdminPhone:    data["admin_phone"].(string),
		Permission:    data["admin_permission"].(int),
	}
	err = db.Create(&admin).Error
	return
}

func UpdateAdminByUserID(adminID string, fieldName string, it interface{}) (err error) {
	db := global.DB
	err = db.Model(&Admin{}).Where("admin_id = ?", adminID).Update(fieldName, it).Error
	return err
}

func UpdateAdminByAdmin(adminID string, admin Admin) (err error) {
	db := global.DB
	err = db.Model(&Admin{}).Where("admin_id = ?", adminID).Updates(admin).Error
	return
}
