package model

import (
	"database/sql"
	"finders-server/global"
	"github.com/jinzhu/gorm"
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
//type Comment struct {
//	CommentID  int    `gorm:"column:comment_id;type:INT;primary_key" json:"comment_id"`    //[ 0] comment_id                                     INT                  null: false  primary: true   auto: false
//	ItemID string `gorm:"column:activity_id;type:VARCHAR;size:30;" json:"activity_id"` //[ 1] activity_id                                    VARCHAR[30]          null: true   primary: false  auto: false
//	Content    string `gorm:"column:content;type:TEXT;size:65535;" json:"content"`         //[ 3] content                                        TEXT[65535]          null: true   primary: false  auto: false
//	FromUID    string `gorm:"column:from_uid;type:VARCHAR;size:30;" json:"from_uid"`       //[ 4] from_uid                                       VARCHAR[30]          null: true   primary: false  auto: false
//	Status     int    `gorm:"column:status;type:INT;" json:"status"`                       //[ 6] status                                         INT                  null: true   primary: false  auto: false
//	TimeModel
//}

type Comment struct {
	CommentID int    `gorm:"column:comment_id;type:INT;primary_key" json:"comment_id"` //[ 0] comment_id                                     INT                  null: false  primary: true   auto: false
	ItemID    string `gorm:"column:item_id;type:varchar(50);" json:"item_id"`          //[ 1] activity_id                                    VARCHAR[30]          null: true   primary: false  auto: false
	ItemType  int    `gorm:"column:item_type;type:int;" json:"item_type"`              //[ 1] activity_id                                    VARCHAR[30]          null: true   primary: false  auto: false
	Content   string `gorm:"column:content;type:TEXT;size:65535;" json:"content"`      //[ 3] content                                        TEXT[65535]          null: true   primary: false  auto: false
	FromUID   string `gorm:"column:from_uid;type:varchar(50);" json:"from_uid"`        //[ 4] from_uid                                       VARCHAR[30]          null: true   primary: false  auto: false
	ToUID     string `gorm:"column:to_uid;type:varchar(50);" json:"from_uid"`          //[ 4] from_uid                                       VARCHAR[30]          null: true   primary: false  auto: false
	Status    int    `gorm:"column:status;type:INT;" json:"status"`                    //[ 6] status                                         INT                  null: true   primary: false  auto: false
	FromUser  User   `gorm:"foreignkey:user_id;association_foreignkey:from_uid"`       //[ 6] status                                         INT                  null: true   primary: false  auto: false
	TimeModel
}

// item_type
const (
	CommentOnActivity = baseIndex + iota
	CommentOnComment
)

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

// status
const (
	CommentNormal = baseIndex + iota
)

func ExistCommentByMap(data map[string]interface{}) bool {
	db := global.DB
	var comment Comment
	err := db.Where(data).First(&comment).Error
	return !gorm.IsRecordNotFoundError(err)
}

func AddCommentByMap(data map[string]interface{}) (comment Comment, err error) {
	db := global.DB
	comment = Comment{
		ItemID:   data["item_id"].(string),
		ItemType: data["item_type"].(int),
		Content:  data["content"].(string),
		FromUID:  data["from_uid"].(string),
		ToUID:    data["to_uid"].(string),
		Status:   CommentNormal,
	}
	err = db.Create(&comment).Error
	return
}

func GetCommentByCommentID(commentID int) (comment Comment, err error) {
	db := global.DB
	err = db.Where("comment_id = ?", commentID).First(&comment).Error
	return
}

func GetCommentTotal(itemID string, itemType int) (cnt int, err error) {
	db := global.DB
	err = db.Model(&Comment{}).Where("item_id = ? AND item_type = ?", itemID, itemType).Count(&cnt).Error
	return
}

func GetCommentsByItemID(pageNum, pageSize int, itemID string, itemType int) (comments []*Comment, err error) {
	db := global.DB
	err = db.Preload("FromUser").Where("item_id = ? AND item_type = ?", itemID, itemType).Offset(pageNum).Limit(pageSize).Find(&comments).Error
	return
}
