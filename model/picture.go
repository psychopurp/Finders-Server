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


CREATE TABLE `pictures` (
  `picture_id` varchar(30) NOT NULL COMMENT '图片ID',
  `picture_url` varchar(200) DEFAULT NULL COMMENT '图片地址',
  `picture_type` int(11) DEFAULT NULL COMMENT '图片类型',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `user_id` varchar(30) NOT NULL COMMENT '用户ID',
  PRIMARY KEY (`picture_id`),
  KEY `fk_picture_user_id` (`user_id`),
  CONSTRAINT `fk_picture_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Picture struct is a row record of the pictures table in the employees database
type Picture struct {
	PictureID   string      `gorm:"column:picture_id;type:VARCHAR;size:30;primary_key" json:"picture_id"` //[ 0] picture_id                                     VARCHAR[30]          null: false  primary: true   auto: false
	PictureURL  null.String `gorm:"column:picture_url;type:VARCHAR;size:200;" json:"picture_url"`         //[ 1] picture_url                                    VARCHAR[200]         null: true   primary: false  auto: false
	PictureType null.Int    `gorm:"column:picture_type;type:INT;" json:"picture_type"`                    //[ 2] picture_type                                   INT                  null: true   primary: false  auto: false
	CreatedAt   time.Time   `gorm:"column:created_at;type:DATETIME;" json:"created_at"`                   //[ 3] created_at                                     DATETIME             null: false  primary: false  auto: false
	UserID      string      `gorm:"column:user_id;type:VARCHAR;size:30;" json:"user_id"`                  //[ 4] user_id                                        VARCHAR[30]          null: false  primary: false  auto: false

}

// TableName sets the insert table name for this struct type
func (p *Picture) TableName() string {
	return "pictures"
}

func (p *Picture) BeforeSave() error {
	return nil
}

func (p *Picture) Prepare() {
}

func (p *Picture) Validate(action Action) error {

	return nil
}
