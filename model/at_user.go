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


CREATE TABLE `at_users` (
  `atusers_id` int(11) NOT NULL COMMENT '艾特ID',
  `acitvity_id` varchar(30) DEFAULT NULL COMMENT '帖子ID',
  `from_uid` varchar(30) DEFAULT NULL COMMENT '用户ID',
  `to_uid` varchar(30) DEFAULT NULL COMMENT '被艾特用户ID',
  `created_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '艾特时间',
  PRIMARY KEY (`atusers_id`),
  KEY `fk_at_activity` (`acitvity_id`),
  KEY `fk_at_user1` (`from_uid`),
  KEY `fk_at_user2` (`to_uid`),
  CONSTRAINT `fk_at_activity` FOREIGN KEY (`acitvity_id`) REFERENCES `activities` (`activity_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_at_user1` FOREIGN KEY (`from_uid`) REFERENCES `users` (`user_id`),
  CONSTRAINT `fk_at_user2` FOREIGN KEY (`to_uid`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// AtUser struct is a row record of the at_users table in the employees database
type AtUser struct {
	AtusersID  int         `gorm:"column:atusers_id;type:INT;primary_key" json:"atusers_id"`    //[ 0] atusers_id                                     INT                  null: false  primary: true   auto: false
	AcitvityID null.String `gorm:"column:acitvity_id;type:VARCHAR;size:30;" json:"acitvity_id"` //[ 1] acitvity_id                                    VARCHAR[30]          null: true   primary: false  auto: false
	FromUID    null.String `gorm:"column:from_uid;type:VARCHAR;size:30;" json:"from_uid"`       //[ 2] from_uid                                       VARCHAR[30]          null: true   primary: false  auto: false
	ToUID      null.String `gorm:"column:to_uid;type:VARCHAR;size:30;" json:"to_uid"`           //[ 3] to_uid                                         VARCHAR[30]          null: true   primary: false  auto: false
	CreatedAt  null.Time   `gorm:"column:created_at;type:DATETIME;" json:"created_at"`          //[ 4] created_at                                     DATETIME             null: true   primary: false  auto: false

}

// TableName sets the insert table name for this struct type
func (a *AtUser) TableName() string {
	return "at_users"
}

func (a *AtUser) BeforeSave() error {
	return nil
}

func (a *AtUser) Prepare() {
}

func (a *AtUser) Validate(action Action) error {

	return nil
}
