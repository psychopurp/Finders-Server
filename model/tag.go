package model

import (
	"database/sql"
	"finders-server/global"
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


CREATE TABLE `tags` (
  `tag_id` int(11) NOT NULL COMMENT '标签ID',
  `tag_name` varchar(50) NOT NULL COMMENT '标签名',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `tag_type` int(11) DEFAULT NULL COMMENT '标签类型',
  PRIMARY KEY (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Tag struct is a row record of the tags table in the employees database
type Tag struct {
	TagID   int    `gorm:"column:tag_id;type:INT;primary_key" json:"tag_id"`  //[ 0] tag_id                                         INT                  null: false  primary: true   auto: false
	TagName string `gorm:"column:tag_name;type:varchar(50);" json:"tag_name"` //[ 1] tag_name                                       VARCHAR[50]          null: false  primary: false  auto: false
	TagType int    `gorm:"column:tag_type;type:INT;" json:"tag_type"`         //[ 3] tag_type                                       INT                  null: true   primary: false  auto: false
	//CreatedAt time.Time `gorm:"column:created_at;type:DATETIME;" json:"created_at"`    //[ 2] created_at                                     DATETIME             null: false  primary: false  auto: false
	TimeModel
}

// TableName sets the insert table name for this struct type
func (t *Tag) TableName() string {
	return "tags"
}

func (t *Tag) BeforeSave() error {
	return nil
}

func (t *Tag) Prepare() {
}

func (t *Tag) Validate(action Action) error {

	return nil
}

// type
const (
	TagSystem = baseIndex + iota
	TagDIY
)

func GetTagType(tagType int) (typeName string, ok bool) {
	data := map[int]string{
		TagSystem: "system",
		TagDIY:    "diy",
	}
	typeName, ok = data[tagType]
	return
}

func GetTagsByActivityID(activityID string) (tags []*Tag, err error) {
	var (
		ids []int
	)
	db := global.DB
	ids, err = GetTagsIDOnTagMap(activityID, TagActivityType)
	if err != nil {
		return
	}
	err = db.Where("tag_id IN (?)", ids).Find(&tags).Error
	return
}
