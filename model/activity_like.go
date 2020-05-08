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


CREATE TABLE `activity_likes` (
  `activity_id` varchar(30) NOT NULL COMMENT '帖子ID',
  `user_id` varchar(30) NOT NULL COMMENT '用户ID',
  `id` int(11) NOT NULL,
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '点赞时间',
  PRIMARY KEY (`id`),
  KEY `fk_likes_activity` (`activity_id`),
  KEY `fk_likes_user` (`user_id`),
  CONSTRAINT `fk_likes_activity` FOREIGN KEY (`activity_id`) REFERENCES `activities` (`activity_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_likes_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// ActivityLike struct is a row record of the activity_likes table in the employees database
type ActivityLike struct {
	ActivityID string    `gorm:"column:activity_id;type:VARCHAR;size:30;" json:"activity_id"` //[ 0] activity_id                                    VARCHAR[30]          null: false  primary: false  auto: false
	UserID     string    `gorm:"column:user_id;type:VARCHAR;size:30;" json:"user_id"`         //[ 1] user_id                                        VARCHAR[30]          null: false  primary: false  auto: false
	ID         int       `gorm:"column:id;type:INT;primary_key" json:"id"`                    //[ 2] id                                             INT                  null: false  primary: true   auto: false
	CreatedAt  time.Time `gorm:"column:created_at;type:DATETIME;" json:"created_at"`          //[ 3] created_at                                     DATETIME             null: false  primary: false  auto: false

}

// TableName sets the insert table name for this struct type
func (a *ActivityLike) TableName() string {
	return "activity_likes"
}

func (a *ActivityLike) BeforeSave() error {
	return nil
}

func (a *ActivityLike) Prepare() {
}

func (a *ActivityLike) Validate(action Action) error {

	return nil
}
