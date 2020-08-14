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


CREATE TABLE `replies` (
  `reply_id` int(11) NOT NULL COMMENT '回复ID',
  `comment_id` int(11) DEFAULT NULL COMMENT '评论ID',
  `reply_type` int(11) DEFAULT NULL COMMENT '回复类型',
  `to_reply_id` int(11) DEFAULT NULL COMMENT '回复目标ID',
  `content` text,
  `from_uid` varchar(30) DEFAULT NULL COMMENT '回复用户ID',
  `to_uid` varchar(30) DEFAULT NULL COMMENT '目标用户ID',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
  `status` int(11) DEFAULT NULL COMMENT '回复状态',
  PRIMARY KEY (`reply_id`),
  KEY `fk_reply_comment` (`comment_id`),
  KEY `fk_reply_reply` (`to_reply_id`),
  KEY `fk_reply_from_user` (`from_uid`),
  KEY `fk_reply_to_user` (`to_uid`),
  CONSTRAINT `fk_reply_comment` FOREIGN KEY (`comment_id`) REFERENCES `comments` (`comment_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_reply_from_user` FOREIGN KEY (`from_uid`) REFERENCES `users` (`user_id`),
  CONSTRAINT `fk_reply_reply` FOREIGN KEY (`to_reply_id`) REFERENCES `replies` (`reply_id`),
  CONSTRAINT `fk_reply_to_user` FOREIGN KEY (`to_uid`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Reply struct is a row record of the replies table in the employees database
type Reply struct {
	ReplyID   int         `gorm:"column:reply_id;type:INT;primary_key" json:"reply_id"`  //[ 0] reply_id                                       INT                  null: false  primary: true   auto: false
	CommentID null.Int    `gorm:"column:comment_id;type:INT;" json:"comment_id"`         //[ 1] comment_id                                     INT                  null: true   primary: false  auto: false
	ReplyType null.Int    `gorm:"column:reply_type;type:INT;" json:"reply_type"`         //[ 2] reply_type                                     INT                  null: true   primary: false  auto: false
	ToReplyID null.Int    `gorm:"column:to_reply_id;type:INT;" json:"to_reply_id"`       //[ 3] to_reply_id                                    INT                  null: true   primary: false  auto: false
	Content   null.String `gorm:"column:content;type:TEXT;size:65535;" json:"content"`   //[ 4] content                                        TEXT[65535]          null: true   primary: false  auto: false
	FromUID   null.String `gorm:"column:from_uid;type:VARCHAR;size:30;" json:"from_uid"` //[ 5] from_uid                                       VARCHAR[30]          null: true   primary: false  auto: false
	ToUID     null.String `gorm:"column:to_uid;type:VARCHAR;size:30;" json:"to_uid"`     //[ 6] to_uid                                         VARCHAR[30]          null: true   primary: false  auto: false
	Status    null.Int    `gorm:"column:status;type:INT;" json:"status"`                 //[ 9] status                                         INT                  null: true   primary: false  auto: false
	//CreatedAt time.Time   `gorm:"column:created_at;type:DATETIME;" json:"created_at"`    //[ 7] created_at                                     DATETIME             null: false  primary: false  auto: false
	//DeletedAt null.Time   `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`    //[ 8] deleted_at                                     DATETIME             null: true   primary: false  auto: false
	TimeModel
}

// TableName sets the insert table name for this struct type
func (r *Reply) TableName() string {
	return "replies"
}

func (r *Reply) BeforeSave() error {
	return nil
}

func (r *Reply) Prepare() {
}

func (r *Reply) Validate(action Action) error {

	return nil
}
