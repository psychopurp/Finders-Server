package model

import (
<<<<<<< HEAD
=======
	"finders-server/global"
>>>>>>> test
	"time"
)

/*
DB Table Details
登陆信息表

CREATE TABLE `logins` (
  `login_id` int(11) NOT NULL COMMENT '登陆ID',
  `user_id` varchar(30) NOT NULL COMMENT '用户ID',
  `access_token` varchar(200) NOT NULL COMMENT 'api token',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '登陆时间',
  `expired_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '过期时间',
  PRIMARY KEY (`login_id`),
  KEY `fk_login_user` (`user_id`),
  CONSTRAINT `fk_login_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Login struct is a row record of the logins table in the employees database
type Login struct {
	LoginID     int       `gorm:"column:login_id;type:INT;primary_key" json:"login_id"`           //[ 0] login_id                                       INT                  null: false  primary: true   auto: false
<<<<<<< HEAD
	UserID      string    `gorm:"column:user_id;type:VARCHAR;size:30;" json:"user_id"`            //[ 1] user_id                                        VARCHAR[30]          null: false  primary: false  auto: false
	AccessToken string    `gorm:"column:access_token;type:VARCHAR;size:200;" json:"access_token"` //[ 2] access_token                                   VARCHAR[200]         null: false  primary: false  auto: false
	CreatedAt   time.Time `gorm:"column:created_at;type:DATETIME;" json:"created_at"`             //[ 3] created_at                                     DATETIME             null: false  primary: false  auto: false
	ExpiredAt   time.Time `gorm:"column:expired_at;type:DATETIME;" json:"expired_at"`             //[ 4] expired_at                                     DATETIME             null: false  primary: false  auto: false

=======
	UserID      string    `gorm:"column:user_id;type:VARCHAR;size:50;" json:"user_id"`            //[ 1] user_id                                        VARCHAR[30]          null: false  primary: false  auto: false
	AccessToken string    `gorm:"column:access_token;type:VARCHAR;size:250;" json:"access_token"` //[ 2] access_token                                   VARCHAR[200]         null: false  primary: false  auto: false
	CreatedAt   time.Time `gorm:"column:created_at;type:DATETIME;" json:"created_at"`             //[ 3] created_at                                     DATETIME             null: false  primary: false  auto: false
	ExpiredAt   time.Time `gorm:"column:expired_at;type:DATETIME;" json:"expired_at"`             //[ 4] expired_at                                     DATETIME             null: false  primary: false  auto: false
>>>>>>> test
}

// TableName sets the insert table name for this struct type
func (l *Login) TableName() string {
	return "logins"
}

func (l *Login) BeforeSave() error {
	return nil
}

func (l *Login) Prepare() {
}

func (l *Login) Validate(action Action) error {

	return nil
}
<<<<<<< HEAD
=======

func AddLoginRecord(userID, token string, createAt, expiredAt time.Time) bool{
	db := global.DB
	db.Create(&Login{
		UserID:      userID,
		AccessToken: token,
		CreatedAt:   createAt,
		ExpiredAt:   expiredAt,
	})
	return true
}
>>>>>>> test
