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


CREATE TABLE `communities` (
  `community_id` int(11) NOT NULL COMMENT '社区ID',
  `community_creator` varchar(30) NOT NULL COMMENT '社区创建者（圈主）',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `community_name` varchar(100) NOT NULL COMMENT '社区名称',
  `community_description` text,
  `community_status` int(11) NOT NULL COMMENT '社区状态',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `background` varchar(200) DEFAULT NULL COMMENT '背景',
  `board` text,
  `deleted_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
  PRIMARY KEY (`community_id`),
  KEY `fk_com_user` (`community_creator`),
  CONSTRAINT `fk_com_user` FOREIGN KEY (`community_creator`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Community struct is a row record of the communities table in the employees database
type Community struct {
	CommunityID          int         `gorm:"column:community_id;type:INT;primary_key" json:"community_id"`                    //[ 0] community_id                                   INT                  null: false  primary: true   auto: false
	CommunityCreator     string      `gorm:"column:community_creator;type:VARCHAR;size:30;" json:"community_creator"`         //[ 1] community_creator                              VARCHAR[30]          null: false  primary: false  auto: false
	CreatedAt            time.Time   `gorm:"column:created_at;type:DATETIME;" json:"created_at"`                              //[ 2] created_at                                     DATETIME             null: false  primary: false  auto: false
	CommunityName        string      `gorm:"column:community_name;type:VARCHAR;size:100;" json:"community_name"`              //[ 3] community_name                                 VARCHAR[100]         null: false  primary: false  auto: false
	CommunityDescription null.String `gorm:"column:community_description;type:TEXT;size:65535;" json:"community_description"` //[ 4] community_description                          TEXT[65535]          null: true   primary: false  auto: false
	CommunityStatus      int         `gorm:"column:community_status;type:INT;" json:"community_status"`                       //[ 5] community_status                               INT                  null: false  primary: false  auto: false
	UpdatedAt            null.Time   `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"`                              //[ 6] updated_at                                     DATETIME             null: true   primary: false  auto: false
	Background           null.String `gorm:"column:background;type:VARCHAR;size:200;" json:"background"`                      //[ 7] background                                     VARCHAR[200]         null: true   primary: false  auto: false
	Board                null.String `gorm:"column:board;type:TEXT;size:65535;" json:"board"`                                 //[ 8] board                                          TEXT[65535]          null: true   primary: false  auto: false
	DeletedAt            null.Time   `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`                              //[ 9] deleted_at                                     DATETIME             null: true   primary: false  auto: false

}

// TableName sets the insert table name for this struct type
func (c *Community) TableName() string {
	return "communities"
}

func (c *Community) BeforeSave() error {
	return nil
}

func (c *Community) Prepare() {
}

func (c *Community) Validate(action Action) error {

	return nil
}
