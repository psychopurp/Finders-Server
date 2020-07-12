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

const (
	USER      = 1
	ACTIVITY  = 2
	COMMUNITY = 3
	MEDIA     = 4
)

type TagMap struct {
	ItemID    string     `gorm:"column:item_id;type:varchar(100);primary_key" json:"item_id"` //类型对应的ID
	ItemType  int        `gorm:"column:item_type;type:INT;" json:"item_type"`
	TagID     int        `gorm:"column:tag_id;type:INT;primary_key" json:"tag_id"`   //[ 0] tag_id                                         INT                  null: false  primary: true   auto: false
	CreatedAt time.Time  `gorm:"column:created_at;type:DATETIME;" json:"created_at"` //[ 2] created_at                                     DATETIME             null: false  primary: false  auto: false
	DeletedAt *time.time `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`
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
