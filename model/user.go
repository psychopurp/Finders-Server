package model

import (
	"finders-server/global"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

const (
	Normal   = 0 //正常状态
	Blocked  = 1 //封号状态
	Canceled = 2 //注销状态
)

/*
DB Table Details
用户基本信息表

CREATE TABLE `users` (
  `user_id` varchar(50) NOT NULL COMMENT '用户ID',
  `phone` varchar(30) NOT NULL COMMENT '手机号',
  `password` varchar(30) NOT NULL COMMENT '密码',
  `nickname` varchar(30) NOT NULL COMMENT '昵称',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '注册时间',
  `status` int(11) NOT NULL COMMENT '用户状态',
  `deleted_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
  `avatar` varchar(100) NOT NULL COMMENT '用户头像',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `unique_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// User struct is a row record of the users table in the  database
type User struct {
	UserID    uuid.UUID  `gorm:"column:user_id;type:varchar(50);primary_key" json:"user_id"`           //[ 0] user_id                                        VARCHAR[30]          null: false  primary: true   auto: false
	Phone     string     `gorm:"column:phone;type:varchar(30);unique_index:unique_phone" json:"phone"` //[ 1] phone                                          VARCHAR[30]          null: false  primary: false  auto: false
	Password  string     `gorm:"column:password;type:varchar(100);" json:"password"`                   //[ 2] password                                       VARCHAR[30]          null: false  primary: false  auto: false
	Nickname  string     `gorm:"column:nickname;type:varchar(30);" json:"nickname"`                    //[ 3] nickname                                       VARCHAR[30]          null: false  primary: false  auto: false
	Status    int        `gorm:"column:status;type:INT;" json:"status"`                                //[ 5] status                                         INT                  null: false  primary: false  auto: false
	CreatedAt time.Time  `gorm:"column:created_at;type:DATETIME;" json:"created_at"`                   //[ 4] created_at                                     DATETIME             null: false  primary: false  auto: false
	UpdatedAt time.Time  `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"`                   //[16] updated_at                                     DATETIME             strue   primary: false  auto: false
	DeletedAt *time.Time `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`                   //[ 6] deleted_at                                     DATETIME             null: true   primary: false  auto: false
	Avatar    string     `gorm:"column:avatar;type:varchar(100);" json:"avatar"`                       //[ 7] avatar                                         VARCHAR[100]         null: false  primary: false  auto: false
	UserInfo  UserInfo   `gorm:"foreignkey:UserId"`                                                    //一对一关系
	UserName  string     `gorm:"column:username;type:varchar(50);unique_index:unique_username" json:"userName"`
	//Relations []Relation `gorm:"many2many:relations;foreignkey:from_uid;association_jointable_foreignkey:relation_id"` //多对多关系
	Relations []Relation `gorm:"many2many:relations;"` //多对多关系
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeSave() error {

	return nil
}

func (u *User) Prepare() {
	fmt.Println("User Prepare()..")
}

func (u *User) Validate(action Action) (err error) {
	err = nil
	switch action {
	case Create:
		return
	case Read:
		return
	case Update:
		return
	case Delete:
		return
	default:
		return
	}
}

// 创建用户后 创建用户信息表
func (u *User) AfterCreate(scope *gorm.Scope) error {
	uinfo := new(UserInfo)
	uinfo.UserID = u.UserID
	uinfo.TrueName = "rotob"
	u.UserInfo = *uinfo

	if err := scope.DB().Create(uinfo).Error; err != nil {
		return err
	} else {
		return nil
	}

}

// func (u User) String() string {
// 	return "this is "
// }

func AddUser(user *User) (err error) {
	db := global.DB
	err = db.Create(user).Error
	return
}

func AddUserByMap(data map[string]interface{}) (user User, err error) {
	db := global.DB
	user = User{
		UserID:   data["user_id"].(uuid.UUID),
		Phone:    data["phone"].(string),
		Password: data["password"].(string),
		Nickname: data["nickname"].(string),
		Status:   data["status"].(int),
		UserName: data["username"].(string),
	}
	err = db.Create(&user).Error
	return
}

func ExistUserByUserNameAndPassword(name, password string) (user User, isExist bool) {
	db := global.DB
	data := make(map[string]interface{})
	data["username"] = name
	data["password"] = password
	data["status"] = Normal
	isExist = !db.Where(data).First(&user).RecordNotFound()
	return
}

func ExistUserByPhone(phone string) (user User, isExist bool) {
	db := global.DB
	data := make(map[string]interface{})
	data["phone"] = phone
	data["status"] = Normal
	isExist = !db.Where(data).First(&user).RecordNotFound()
	return
}

func GetUserByUserName(userName string) (user User, err error) {
	db := global.DB
	err = db.Where("username = ?", userName).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	err = db.Model(&user).Related(&user.UserInfo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	return
}

func GetUserByUserID(userID string) (user User, err error) {
	db := global.DB
	err = db.Where("user_id = ?", userID).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	err = db.Model(&user).Related(&user.UserInfo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	return
}

func GetUsersByUserIDs(userIDs []string) (users []*User, err error) {
	db := global.DB
	//err = db.Where("user_id IN (?)", userIDs).Find(&users).Error
	err = db.Preload("UserInfo").Where("user_id IN (?)", userIDs).Find(&users).Error
	return
}

func UpdateUserByUserID(userID string, fieldName string, it interface{}) (err error) {
	var user User
	db := global.DB
	err = db.Model(&user).Where("user_id = ?", userID).Update(fieldName, it).Error
	return err
}

func UpdateUserByUser(userID string, user User) (err error) {
	db := global.DB
	err = db.Model(&User{}).Where("user_id = ?", userID).Updates(user).Error
	return
}
func GetUsers() (users []*User, err error) {
	db := global.DB
	err = db.Model(&User{}).Find(&users).Error
	return
}
