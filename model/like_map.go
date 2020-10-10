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
type LikeMap struct {
	ID         int    `gorm:"column:id;type:INT;primary_key" json:"id"`            //[ 2] id                                             INT                  null: false  primary: true   auto: false
	ObjectID   string `gorm:"column:object_id;type:varchar(50);" json:"object_id"` //[ 0] activity_id                                    VARCHAR[30]          null: false  primary: false  auto: false
	ObjectType int    `gorm:"column:object_type;type:int;" json:"object_type"`
	UserID     string `gorm:"column:user_id;type:varchar(50);" json:"user_id"` //[ 1] user_id                                        VARCHAR[30]          null: false  primary: false  auto: false
	TimeModel
}

// TableName sets the insert table name for this struct type
func (a *LikeMap) TableName() string {
	return "like_maps"
}

func (a *LikeMap) BeforeSave() error {
	return nil
}

func (a *LikeMap) Prepare() {
}

func (a *LikeMap) Validate(action Action) error {

	return nil
}

// type
const (
	LikeActivity = baseIndex + iota
	LikeMoment
	LikeQuestionBox
)

func ExistLikeMap(objectID, userID string, objectType int) bool {
	db := global.DB
	var likeMap LikeMap
	err := db.Where("object_id = ? AND user_id = ? AND object_type = ?", objectID, userID, objectType).First(&likeMap).Error
	return !gorm.IsRecordNotFoundError(err)
}

func AddLikeMap(objectID, userID string, objectType int) (likeMap LikeMap, err error) {
	db := global.DB
	likeMap = LikeMap{
		ObjectID:   objectID,
		ObjectType: objectType,
		UserID:     userID,
	}
	err = db.Create(&likeMap).Error
	return
}

func (a *AffairService) AddLikeMap(objectID, userID string, objectType int) (err error) {
	db := a.tx
	likeMap := &LikeMap{
		ObjectID:   objectID,
		ObjectType: objectType,
		UserID:     userID,
	}
	err = db.Model(&LikeMap{}).Create(likeMap).Error
	return
}

func DeleteLikeMap(objectID, userID string, objectType int) (err error) {
	db := global.DB
	err = db.Where("object_id = ? AND user_id = ? AND object_type = ?", objectID, userID, objectType).Delete(&LikeMap{}).Error
	return
}

func GetUserLikeMapTotal(userID string, objectType int) (cnt int, err error) {
	db := global.DB
	err = db.Model(&LikeMap{}).Where("user_id = ? AND object_type = ?", userID, objectType).Count(&cnt).Error
	return
}

func GetLikeMapNumByObjectID(objectID string) (cnt int, err error) {
	db := global.DB
	err = db.Model(&LikeMap{}).Where("object_id = ?", objectID).Count(&cnt).Error
	return
}

func GetLikeMapsByUserID(userID string, objectTYpe int) (likeMaps []*LikeMap, err error) {
	db := global.DB
	err = db.Where("user_id = ? AND object_type = ?", userID, objectTYpe).Find(&likeMaps).Error
	return
}
