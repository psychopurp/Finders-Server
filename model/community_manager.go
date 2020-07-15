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


CREATE TABLE `community_managers` (
  `community_id` int(11) NOT NULL COMMENT '社区ID',
  `id` int(11) NOT NULL COMMENT 'ID',
  `manager_id` varchar(30) DEFAULT NULL COMMENT '管理员ID',
  `permission` int(11) DEFAULT NULL COMMENT '圈子管理员权限',
  `status` int(11) NOT NULL COMMENT '管理员状态',
  PRIMARY KEY (`id`),
  KEY `fk_manager_community` (`community_id`),
  KEY `fk_manager_user` (`manager_id`),
  CONSTRAINT `fk_manager_community` FOREIGN KEY (`community_id`) REFERENCES `communities` (`community_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_manager_user` FOREIGN KEY (`manager_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// CommunityManager struct is a row record of the community_managers table in the employees database
type CommunityManager struct {
	CommunityID int    `gorm:"column:community_id;type:INT;" json:"community_id"`     //[ 0] community_id                                   INT                  null: false  primary: false  auto: false
	ID          int    `gorm:"column:id;type:INT;primary_key" json:"id"`              //[ 1] id                                             INT                  null: false  primary: true   auto: false
	ManagerID   string `gorm:"column:manager_id;type:varchar(50);" json:"manager_id"` //[ 2] manager_id                                     VARCHAR[30]          null: true   primary: false  auto: false
	Permission  int    `gorm:"column:permission;type:INT;" json:"permission"`         //[ 3] permission                                     INT                  null: true   primary: false  auto: false
	Status      int    `gorm:"column:status;type:INT;" json:"status"`                 //[ 4] status                                         INT                  null: false  primary: false  auto: false
	TimeModel
}

// TableName sets the insert table name for this struct type
func (c *CommunityManager) TableName() string {
	return "community_managers"
}

func (c *CommunityManager) BeforeSave() error {
	return nil
}

func (c *CommunityManager) Prepare() {
}

func (c *CommunityManager) Validate(action Action) error {

	return nil
}
