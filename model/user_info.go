package model

import (
	"finders-server/global"
	"time"

	uuid "github.com/satori/go.uuid"
)

/*
DB Table Details
用户详细信息表

CREATE TABLE `user_infos` (
  `user_id` varchar(50) NOT sOMMENT '用户ID',
  `truename` varchar(40) DEFAULT sOMMENT '真实姓名',
  `address` varchar(200) DEFAULT sOMMENT '所在地',
  `sex` varchar(4) DEFAULT sOMMENT '性别',
  `sexual` varchar(8) DEFAULT sOMMENT '性取向',
  `feeling` varchar(20) DEFAULT sOMMENT '感情状况',
  `birthday` varchar(20) DEFAULT sOMMENT '生日',
  `introduction` varchar(400) DEFAULT sOMMENT '简介',
  `blood_type` varchar(8) DEFAULT sOMMENT '血型',
  `eamil` varchar(60) DEFAULT sOMMENT '邮箱',
  `qq` varchar(30) DEFAULT sOMMENT 'QQ',
  `wechat` varchar(30) DEFAULT sOMMENT '微信',
  `profession` varchar(60) DEFAULT sOMMENT '职业信息',
  `school` varchar(30) DEFAULT sOMMENT '学校',
  `constellation` varchar(40) DEFAULT sOMMENT '星座',
  `created_at` datetime NOT sN UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT sN UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `credit` int(11) NOT sOMMENT '用户信誉积分',
  `user_tag` text,
  `deleted_at` datetime DEFAULT sN UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// UserInfo struct is a row record of the user_infos table in the  database
type UserInfo struct {
	UserID        uuid.UUID  `gorm:"column:user_id;type:varchar(50);primary_key;" json:"user_id"` //[ 0] user_id                                        VARCHAR[30]          sfalse  primary: true   auto: false
	TrueName      string     `gorm:"column:truename;type:varchar(40);" json:"truename"`           //[ 1] truename                                       VARCHAR[40]          strue   primary: false  auto: false
	Address       string     `gorm:"column:address;type:varchar(200);" json:"address"`            //[ 2] address                                        VARCHAR[200]         strue   primary: false  auto: false
	Sex           string     `gorm:"column:sex;type:varchar(4);" json:"sex"`                      //[ 3] sex                                            VARCHAR[4]           strue   primary: false  auto: false
	Sexual        string     `gorm:"column:sexual;type:varchar(8);" json:"sexual"`                //[ 4] sexual                                         VARCHAR[8]           strue   primary: false  auto: false
	Feeling       string     `gorm:"column:feeling;type:varchar(20);" json:"feeling"`             //[ 5] feeling                                        VARCHAR[20]          strue   primary: false  auto: false
	Birthday      string     `gorm:"column:birthday;type:varchar(20);" json:"birthday"`           //[ 6] birthday                                       VARCHAR[20]          strue   primary: false  auto: false
	Introduction  string     `gorm:"column:introduction;type:varchar(400);" json:"introduction"`  //[ 7] introduction                                   VARCHAR[400]         strue   primary: false  auto: false
	BloodType     string     `gorm:"column:blood_type;type:varchar(8);" json:"blood_type"`        //[ 8] blood_type                                     VARCHAR[8]           strue   primary: false  auto: false
	Eamil         string     `gorm:"column:eamil;type:varchar(60);" json:"eamil"`                 //[ 9] eamil                                          VARCHAR[60]          strue   primary: false  auto: false
	QQ            string     `gorm:"column:qq;type:varchar(30);" json:"qq"`                       //[10] qq                                             VARCHAR[30]          strue   primary: false  auto: false
	Wechat        string     `gorm:"column:wechat;type:varchar(30);" json:"wechat"`               //[11] wechat                                         VARCHAR[30]          strue   primary: false  auto: false
	Profession    string     `gorm:"column:profession;type:varchar(60);" json:"profession"`       //[12] profession                                     VARCHAR[60]          strue   primary: false  auto: false
	School        string     `gorm:"column:school;type:varchar(30);" json:"school"`               //[13] school                                         VARCHAR[30]          strue   primary: false  auto: false
	Constellation string     `gorm:"column:constellation;type:varchar(40);" json:"constellation"` //[14] constellation                                  VARCHAR[40]          strue   primary: false  auto: false
	CreatedAt     time.Time  `gorm:"column:created_at;type:DATETIME;" json:"created_at"`          //[15] created_at                                     DATETIME             sfalse  primary: false  auto: false
	UpdatedAt     time.Time  `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"`          //[16] updated_at                                     DATETIME             strue   primary: false  auto: false
	Credit        int        `gorm:"column:credit;type:INT;" json:"credit"`                       //[17] credit                                         INT                  sfalse  primary: false  auto: false
	UserTag       string     `gorm:"column:user_tag;type:TEXT;size:65535;" json:"user_tag"`       //[18] user_tag                                       TEXT[65535]          strue   primary: false  auto: false
	DeletedAt     *time.Time `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`          //[19] deleted_at                                     DATETIME             strue   primary: false  auto: false
	Age           int        `gorm:"column:age;type:INT;" json:"age"`                             //[20] age
}

// TableName sets the insert table name for this struct type
func (u *UserInfo) TableName() string {
	return "user_infos"
}

func (u *UserInfo) BeforeSave() error {
	return nil
}

func (u *UserInfo) Prepare() {
}

func (u *UserInfo) Validate(action Action) error {

	return nil
}

func AddUserInfo(userInfo *UserInfo) (err error){
	db := global.DB
	err = db.Create(userInfo).Error
	return
}

func GetUserInfoByUserID(userID string)(userInfo UserInfo, err error){
	db := global.DB
	err = db.Where("user_id = ?", userID).First(&userInfo).Error
	return
}

func UpdateUserInfoByUserID(userID string, fieldName string, it interface{})(err error){
	var userInfo UserInfo
	db := global.DB
	err = db.Model(&userInfo).Where("user_id = ?", userID).Update(fieldName, it).Error
	return err
}

func UpdateUserInfoByUserInfo(userInfo UserInfo) (err error){
	db := global.DB
	err = db.Save(&userInfo).Error
	return
}