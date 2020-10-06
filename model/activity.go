package model

import (
	"database/sql"
	"errors"
	"finders-server/global"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
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
	ActivityID     string `gorm:"column:activity_id;type:varchar(50);primary_key" json:"activity_id"` //[ 0] activity_id                                    VARCHAR[30]          null: false  primary: true   auto: false
	ActivityStatus int    `gorm:"column:activity_status;type:INT;" json:"activity_status"`            //[ 1] activity_status                                INT                  null: true   primary: false  auto: false
	ActivityInfo   string `gorm:"column:activity_info;type:TEXT;size:65535;" json:"activity_info"`    //[ 2] activity_info                                  TEXT[65535]          null: true   primary: false  auto: false
	ActivityTitle  string `gorm:"column:activity_title;type:varchar(50);" json:"activity_title"`      //[ 2] activity_info                                  TEXT[65535]          null: true   primary: false  auto: false
	CollectNum     int    `gorm:"column:collect_num;type:INT;" json:"collect_num"`                    //[ 4] collect_num                                    INT                  null: false  primary: false  auto: false
	CommentNum     int    `gorm:"column:comment_num;type:INT;" json:"comment_num"`                    //[ 5] comment_num                                    INT                  null: false  primary: false  auto: false
	ReadNum        int    `gorm:"column:read_num;type:INT;" json:"read_num"`                          //[ 6] read_num                                       INT                  null: false  primary: false  auto: false
	//ActivityTag    string `gorm:"column:activity_tag;type:TEXT;size:65535;" json:"activity_tag"`          //[ 7] activity_tag                                   TEXT[65535]          null: true   primary: false  auto: false
	MediaIDs string `gorm:"column:media_ids;varchar(5000);" json:"media_id"` //[ 8] picture_id                                     VARCHAR[30]          null: true   primary: false  auto: false
	//MediaType   int    `gorm:"column:media_type;int;" json:"media_id"`       //[ 8] picture_id                                     VARCHAR[30]          null: true   primary: false  auto: false
	Media       Media  `gorm:"foreignkey:media_id;association_foreignkey:media_id"`
	User        User   `gorm:"foreignkey:user_id;association_foreignkey:user_id"`
	UserID      string `gorm:"column:user_id;type:varchar(50);" json:"user_id"`   //[ 9] user_id                                        VARCHAR[30]          null: true   primary: false  auto: false
	CommunityID int    `gorm:"column:community_id;type:INT;" json:"community_id"` //[10] community_id                                   INT                  null: false  primary: false  auto: false
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

// status
const (
	ActivityNormal = baseIndex + iota
)

func AddActivityByMap(data map[string]interface{}) (activity Activity, err error) {
	db := global.DB
	activity = Activity{
		ActivityID:     uuid.NewV4().String(),
		ActivityStatus: ActivityNormal,
		ActivityInfo:   data["activity_info"].(string),
		ActivityTitle:  data["activity_title"].(string),
		MediaIDs:       data["media_ids"].(string),
		UserID:         data["user_id"].(string),
		CommunityID:    data["community_id"].(int),
	}
	err = db.Create(&activity).Error
	return
}

func (a *AffairService) AddActivityByMap(data map[string]interface{}) (activity Activity, err error) {
	db := a.tx
	activity = Activity{
		ActivityID:     uuid.NewV4().String(),
		ActivityStatus: ActivityNormal,
		ActivityInfo:   data["activity_info"].(string),
		ActivityTitle:  data["activity_title"].(string),
		MediaIDs:       data["media_ids"].(string),
		UserID:         data["user_id"].(string),
		CommunityID:    data["community_id"].(int),
	}
	err = db.Create(&activity).Error
	return
}

func ExistActivityByMap(data map[string]interface{}) bool {
	db := global.DB
	var activity Activity
	err := db.Where(data).First(&activity).Error
	return !gorm.IsRecordNotFoundError(err)
}

func GetActivityByID(activityID string) (activity Activity, err error) {
	db := global.DB
	err = db.Where("activity_id = ?", activityID).First(&activity).Error
	return
}

func GetActivities() (activities []*Activity, err error) {
	db := global.DB
	err = db.Model(&Activity{}).Find(&activities).Error
	return
}
func GetActivitiesByCommunityID(pageNum, pageSize, communityID int) (activities []*Activity, err error) {
	db := global.DB
	err = db.Preload("User").Where("community_id = ?", communityID).Offset(pageNum).Limit(pageSize).Find(&activities).Error
	return
}

func GetActivitiesByUserID(pageNum, pageSize int, userID string) (activities []*Activity, err error) {
	db := global.DB
	err = db.Preload("User").Where("user_id = ?", userID).Offset(pageNum).Limit(pageSize).Find(&activities).Error
	return
}

func GetActivitiesByActivityIDs(pageNum, pageSize int, activityIDs []string) (activities []*Activity, err error) {
	db := global.DB
	err = db.Preload("User").Where("activity_id IN (?)", activityIDs).Offset(pageNum).Limit(pageSize).Find(&activities).Error
	return
}

func GetActivityTotalByCommunityID(communityID int) (cnt int, err error) {
	db := global.DB
	err = db.Model(&Activity{}).Where("community_id = ?", communityID).Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	return
}

func GetActivityTotalByUserID(userID string) (cnt int, err error) {
	db := global.DB
	err = db.Model(&Activity{}).Where("user_id = ?", userID).Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	return
}

const (
	AddOP   = "add"
	MinusOP = "minus"
)

func AddActivityCollectNum(activityID string, op string) (err error) {
	db := global.DB
	var activity Activity
	err = db.Where("activity_id = ?", activityID).First(&activity).Error
	if err != nil {
		return
	}
	if op == AddOP {
		activity.CollectNum = activity.CollectNum + 1
	} else {
		activity.CollectNum = activity.CollectNum - 1
		if activity.CollectNum < 0 {
			return errors.New("no < 0")
		}
	}
	err = db.Save(&activity).Error
	return
}

func (a *AffairService) AddActivityCollectNum(activityID string, op string) (err error) {
	db := a.tx
	var activity Activity
	err = db.Where("activity_id = ?", activityID).First(&activity).Error
	if err != nil {
		return
	}
	if op == AddOP {
		activity.CollectNum = activity.CollectNum + 1
	} else {
		activity.CollectNum = activity.CollectNum - 1
		if activity.CollectNum < 0 {
			return errors.New("no < 0")
		}
	}
	err = db.Save(&activity).Error
	return
}

func UpdateActivityReadNum(activityID string, op string) (err error) {
	db := global.DB
	var activity Activity
	err = db.Where("activity_id = ?", activityID).First(&activity).Error
	if err != nil {
		return
	}
	if op == AddOP {
		activity.ReadNum = activity.ReadNum + 1
	} else {
		activity.ReadNum = activity.ReadNum - 1
		if activity.ReadNum < 0 {
			return errors.New("no < 0")
		}
	}

	err = db.Save(&activity).Error
	return
}

func (a *AffairService) UpdateActivityReadNum(activityID string, op string) (err error) {
	db := a.tx
	var activity Activity
	err = db.Where("activity_id = ?", activityID).First(&activity).Error
	if err != nil {
		return
	}
	if op == AddOP {
		activity.ReadNum = activity.ReadNum + 1
	} else {
		activity.ReadNum = activity.ReadNum - 1
		if activity.ReadNum < 0 {
			return errors.New("no < 0")
		}
	}

	err = db.Save(&activity).Error
	return
}

func UpdateActivityCommentNum(activityID string, op string) (err error) {
	db := global.DB
	var activity Activity
	err = db.Where("activity_id = ?", activityID).First(&activity).Error
	if err != nil {
		return
	}
	if op == AddOP {
		activity.CommentNum = activity.CommentNum + 1
	} else {
		activity.CommentNum = activity.CommentNum - 1
		if activity.CommentNum < 0 {
			return errors.New("no < 0")
		}
	}

	err = db.Save(&activity).Error
	return
}

func (a *AffairService) UpdateActivityCommentNum(activityID string, op string) (err error) {
	db := a.tx
	var activity Activity
	err = db.Where("activity_id = ?", activityID).First(&activity).Error
	if err != nil {
		return
	}
	if op == AddOP {
		activity.CommentNum = activity.CommentNum + 1
	} else {
		activity.CommentNum = activity.CommentNum - 1
		if activity.CommentNum < 0 {
			return errors.New("no < 0")
		}
	}

	err = db.Save(&activity).Error
	return
}
