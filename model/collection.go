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


CREATE TABLE `collections` (
  `collection_id` int(11) NOT NULL COMMENT '收藏ID',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '收藏时间',
  `collection_status` varchar(10) NOT NULL COMMENT '收藏状态',
  `collection_type` int(11) NOT NULL COMMENT '收藏类型',
  `user_id` varchar(30) NOT NULL,
  `link` varchar(100) NOT NULL COMMENT '收藏链接',
  PRIMARY KEY (`collection_id`),
  KEY `fk_collection_user` (`user_id`),
  CONSTRAINT `fk_collection_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Collection struct is a row record of the collections table in the employees database
type Collection struct {
	CollectionID     int    `gorm:"column:collection_id;type:INT;primary_key" json:"collection_id"`          //[ 0] collection_id                                  INT                  null: false  primary: true   auto: false
	CollectionStatus string `gorm:"column:collection_status;type:VARCHAR;size:10;" json:"collection_status"` //[ 2] collection_status                              VARCHAR[10]          null: false  primary: false  auto: false
	CollectionType   int    `gorm:"column:collection_type;type:INT;" json:"collection_type"`                 //[ 3] collection_type                                INT                  null: false  primary: false  auto: false
	UserID           string `gorm:"column:user_id;type:VARCHAR;size:50;" json:"user_id"`                     //[ 4] user_id                                        VARCHAR[30]          null: false  primary: false  auto: false
	Link             string `gorm:"column:link;type:VARCHAR;size:100;" json:"link"`                          //[ 5] link                                           VARCHAR[100]         null: false  primary: false  auto: false
	//CreatedAt        time.Time `gorm:"column:created_at;type:DATETIME;" json:"created_at"`                      //[ 1] created_at                                     DATETIME             null: false  primary: false  auto: false
	TimeModel
}

// TableName sets the insert table name for this struct type
func (c *Collection) TableName() string {
	return "collections"
}

func (c *Collection) BeforeSave() error {
	return nil
}

func (c *Collection) Prepare() {
}

func (c *Collection) Validate(action Action) error {

	return nil
}

// type
const (
	CollectionCommunity = baseIndex + iota
	CollectionActivity
)

// status
const (
	CollectionNormal = "normal"
	CollectionBanned = "banned"
)

func ExistCollectionByMap(data map[string]interface{}) bool {
	db := global.DB
	var collection Collection
	err := db.Where(data).First(&collection).Error
	return !gorm.IsRecordNotFoundError(err)
}

func AddCollectionByMap(data map[string]interface{}) (collection Collection, err error) {
	db := global.DB
	collection = Collection{
		CollectionStatus: CollectionNormal,
		CollectionType:   data["collection_type"].(int),
		UserID:           data["user_id"].(string),
		Link:             data["link"].(string),
	}
	err = db.Create(&collection).Error
	return
}

func DeleteCollectionByMap(data map[string]interface{}) (err error) {
	db := global.DB
	err = db.Where(data).Delete(&Collection{}).Error
	return
}

func GetCollectionIDs(userID string, collectionType int) (collections []*Collection, err error) {
	db := global.DB
	err = db.Model(&Collection{}).Where("user_id = ? AND collection_type = ?", userID, collectionType).Find(&collections).Error
	return
}

func GetCollectionTotal(userID string, collectionType int) (cnt int, err error) {
	db := global.DB
	err = db.Model(&Collection{}).Where("user_id = ? AND collection_type = ?", userID, collectionType).Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	return
}
