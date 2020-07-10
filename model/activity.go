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


CREATE TABLE `activities` (
  `activity_id` varchar(30) NOT NULL COMMENT '帖子ID',
  `activity_status` int(11) DEFAULT NULL COMMENT '帖子类型',
  `activity_info` text,
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `collect_num` int(11) NOT NULL COMMENT '收藏次数',
  `comment_num` int(11) NOT NULL COMMENT '评论次数',
  `read_num` int(11) NOT NULL COMMENT '阅读次数',
  `activity_tag` text COMMENT '帖子标签',
  `picture_id` varchar(30) DEFAULT NULL COMMENT '帖子图片ID',
  `user_id` varchar(30) DEFAULT NULL COMMENT '发表用户ID',
  `community_id` int(11) NOT NULL COMMENT '所属社区ID',
  PRIMARY KEY (`activity_id`),
  KEY `fk_activity_user` (`user_id`),
  KEY `fk_activity_picture` (`picture_id`),
  KEY `fk_activity_community` (`community_id`),
  CONSTRAINT `fk_activity_community` FOREIGN KEY (`community_id`) REFERENCES `communities` (`community_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_activity_picture` FOREIGN KEY (`picture_id`) REFERENCES `pictures` (`picture_id`),
  CONSTRAINT `fk_activity_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Activity struct is a row record of the activities table in the employees database
type Activity struct {
	ActivityID     string      `gorm:"column:activity_id;type:VARCHAR;size:30;primary_key" json:"activity_id"` //[ 0] activity_id                                    VARCHAR[30]          null: false  primary: true   auto: false
	ActivityStatus null.Int    `gorm:"column:activity_status;type:INT;" json:"activity_status"`                //[ 1] activity_status                                INT                  null: true   primary: false  auto: false
	ActivityInfo   null.String `gorm:"column:activity_info;type:TEXT;size:65535;" json:"activity_info"`        //[ 2] activity_info                                  TEXT[65535]          null: true   primary: false  auto: false
	CollectNum     int         `gorm:"column:collect_num;type:INT;" json:"collect_num"`                        //[ 4] collect_num                                    INT                  null: false  primary: false  auto: false
	CommentNum     int         `gorm:"column:comment_num;type:INT;" json:"comment_num"`                        //[ 5] comment_num                                    INT                  null: false  primary: false  auto: false
	ReadNum        int         `gorm:"column:read_num;type:INT;" json:"read_num"`                              //[ 6] read_num                                       INT                  null: false  primary: false  auto: false
	ActivityTag    null.String `gorm:"column:activity_tag;type:TEXT;size:65535;" json:"activity_tag"`          //[ 7] activity_tag                                   TEXT[65535]          null: true   primary: false  auto: false
	PictureID      null.String `gorm:"column:picture_id;type:VARCHAR;size:30;" json:"picture_id"`              //[ 8] picture_id                                     VARCHAR[30]          null: true   primary: false  auto: false
	UserID         null.String `gorm:"column:user_id;type:VARCHAR;size:30;" json:"user_id"`                    //[ 9] user_id                                        VARCHAR[30]          null: true   primary: false  auto: false
	CommunityID    int         `gorm:"column:community_id;type:INT;" json:"community_id"`                      //[10] community_id                                   INT                  null: false  primary: false  auto: false
	//CreatedAt      time.Time   `gorm:"column:created_at;type:DATETIME;" json:"created_at"`                     //[ 3] created_at                                     DATETIME             null: false  primary: false  auto: false
	TimeModel
}

// TableName sets the insert table name for this struct type
func (a *Activity) TableName() string {
	return "activities"
}

func (a *Activity) BeforeSave() error {
	return nil
}

func (a *Activity) Prepare() {
}

func (a *Activity) Validate(action Action) error {

	return nil
}
