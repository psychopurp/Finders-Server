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


CREATE TABLE `community_users` (
  `id` int(11) NOT NULL,
  `user_id` varchar(30) DEFAULT NULL COMMENT '社区用户ID',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '用户加入时间',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `community_id` int(11) DEFAULT NULL COMMENT '社区ID',
  `status` int(11) DEFAULT NULL COMMENT '用户当前在圈子的状态',
  PRIMARY KEY (`id`),
  KEY `fk_community_user_user` (`user_id`),
  KEY `fk_community_user_community` (`community_id`),
  CONSTRAINT `fk_community_user_community` FOREIGN KEY (`community_id`) REFERENCES `communities` (`community_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_community_user_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// CommunityUser struct is a row record of the community_users table in the employees database
type CommunityUser struct {
	ID          int    `gorm:"column:id;type:INT;primary_key" json:"id"`            //[ 0] id                                             INT                  null: false  primary: true   auto: false
	UserID      string `gorm:"column:user_id;type:VARCHAR;size:30;" json:"user_id"` //[ 1] user_id                                        VARCHAR[30]          null: true   primary: false  auto: false
	CommunityID int    `gorm:"column:community_id;type:INT;" json:"community_id"`   //[ 4] community_id                                   INT                  null: true   primary: false  auto: false
	Status      int    `gorm:"column:status;type:INT;" json:"status"`               //[ 5] status                                         INT                  null: true   primary: false  auto: false
	//CreatedAt   time.Time   `gorm:"column:created_at;type:DATETIME;" json:"created_at"`  //[ 2] created_at                                     DATETIME             null: false  primary: false  auto: false
	//UpdatedAt   time.Time   `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"`  //[ 3] updated_at                                     DATETIME             null: false  primary: false  auto: false
	TimeModel
}

// TableName sets the insert table name for this struct type
func (c *CommunityUser) TableName() string {
	return "community_users"
}

func (c *CommunityUser) BeforeSave() error {
	return nil
}

func (c *CommunityUser) Prepare() {
}

func (c *CommunityUser) Validate(action Action) error {

	return nil
}

const (
	CommunityUserWaitForCheck = baseIndex + iota
	CommunityUserNormal
)
