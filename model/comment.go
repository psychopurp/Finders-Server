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


CREATE TABLE `comments` (
  `comment_id` int(11) NOT NULL COMMENT '评论ID',
  `activity_id` varchar(30) DEFAULT NULL COMMENT '评论的帖子ID',
  `activity_type` varchar(100) DEFAULT NULL COMMENT '帖子类型',
  `content` text,
  `from_uid` varchar(30) DEFAULT NULL COMMENT '评论用户ID',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `status` int(11) DEFAULT NULL COMMENT '评论状态',
  `deleted_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
  PRIMARY KEY (`comment_id`),
  KEY `fk_comment_user` (`from_uid`),
  KEY `fk_comment_activity` (`activity_id`),
  CONSTRAINT `fk_comment_activity` FOREIGN KEY (`activity_id`) REFERENCES `activities` (`activity_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_comment_user` FOREIGN KEY (`from_uid`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Comment struct is a row record of the comments table in the employees database
type Comment struct {
	CommentID    int         `gorm:"column:comment_id;type:INT;primary_key" json:"comment_id"`         //[ 0] comment_id                                     INT                  null: false  primary: true   auto: false
	ActivityID   null.String `gorm:"column:activity_id;type:VARCHAR;size:30;" json:"activity_id"`      //[ 1] activity_id                                    VARCHAR[30]          null: true   primary: false  auto: false
	ActivityType null.String `gorm:"column:activity_type;type:VARCHAR;size:100;" json:"activity_type"` //[ 2] activity_type                                  VARCHAR[100]         null: true   primary: false  auto: false
	Content      null.String `gorm:"column:content;type:TEXT;size:65535;" json:"content"`              //[ 3] content                                        TEXT[65535]          null: true   primary: false  auto: false
	FromUID      null.String `gorm:"column:from_uid;type:VARCHAR;size:30;" json:"from_uid"`            //[ 4] from_uid                                       VARCHAR[30]          null: true   primary: false  auto: false
	Status       null.Int    `gorm:"column:status;type:INT;" json:"status"`                            //[ 6] status                                         INT                  null: true   primary: false  auto: false
	//CreatedAt    time.Time   `gorm:"column:created_at;type:DATETIME;" json:"created_at"`               //[ 5] created_at                                     DATETIME             null: false  primary: false  auto: false
	//DeletedAt    null.Time   `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`               //[ 7] deleted_at                                     DATETIME             null: true   primary: false  auto: false
	TimeModel
}

// TableName sets the insert table name for this struct type
func (c *Comment) TableName() string {
	return "comments"
}

func (c *Comment) BeforeSave() error {
	return nil
}

func (c *Comment) Prepare() {
}

func (c *Comment) Validate(action Action) error {

	return nil
}
