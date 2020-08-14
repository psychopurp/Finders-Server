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

const (
	USER      = 1
	ACTIVITY  = 2
	COMMUNITY = 3
	MEDIA     = 4
)

type TagMap struct {
	ItemID   string `gorm:"column:item_id;type:varchar(100);primary_key" json:"item_id"` //类型对应的ID
	ItemType int    `gorm:"column:item_type;type:int;" json:"item_type"`
	TagID    int    `gorm:"column:tag_id;type:INT;primary_key" json:"tag_id"` //[ 0] tag_id                                         INT                  null: false  primary: true   auto: false
	TimeModel
}

// TableName sets the insert table name for this struct type
func (t *TagMap) TableName() string {
	return "tag_map"
}

func (t *TagMap) BeforeSave() error {
	return nil
}

func (t *TagMap) Prepare() {
}

func (t *TagMap) Validate(action Action) error {

	return nil
}

const (
	TagActivityType = baseIndex + iota
)

func GetTagsIDOnTagMap(itemID string, itemType int) (tagIDs []int, err error) {
	db := global.DB
	var tagMaps []TagMap
	err = db.Where("item_id = ? AND item_type = ?", itemID, itemType).Find(&tagMaps).Error
	for _, tagMap := range tagMaps {
		tagIDs = append(tagIDs, tagMap.TagID)
	}
	return
}
