package model

import (
	"database/sql"
	"finders-server/global"
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


CREATE TABLE `pictures` (
  `picture_id` varchar(30) NOT NULL COMMENT '图片ID',
  `picture_url` varchar(200) DEFAULT NULL COMMENT '图片地址',
  `picture_type` int(11) DEFAULT NULL COMMENT '图片类型',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `user_id` varchar(30) NOT NULL COMMENT '用户ID',
  PRIMARY KEY (`picture_id`),
  KEY `fk_picture_user_id` (`user_id`),
  CONSTRAINT `fk_picture_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Picture struct is a row record of the pictures teble in the employees database
type Media struct {
	MediaID   string     `gorm:"column:media_id;type:varchar(50);primary_key" json:"media_id"` //[ 0] picture_id                                     VARCHAR[30]          null: false  primary: true   auto: false
	MediaURL  string     `gorm:"column:media_url;type:varchar(200)" json:"media_url"`          //[ 1] picture_url                                    VARCHAR[200]         null: true   primary: false  auto: false
	MediaType int        `gorm:"column:media_type;type:INT;" json:"media_type"`                //[ 2] picture_type                                   INT                  null: true   primary: false  auto: false
	UserID    string     `gorm:"column:user_id;type:varchar(50)" json:"user_id"`               //[ 4] user_id                                        VARCHAR[30]          null: false  primary: false  auto: false
	CreatedAt time.Time  `gorm:"column:created_at;type:DATETIME;" json:"created_at"`           //[ 3] created_at                                     DATETIME             null: false  primary: false  auto: false
	DeletedAt *time.Time `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`           //[ 6] deleted_at                                     DATETIME             null: true   primary: false  auto: false
	UpdatedAt time.Time  `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"`           //[16] updated_at                                     DATETIME             strue   primary: false  auto: false
}

const (
	PICTURE = baseIndex + iota
	VIDEO
)

func GetMediaTypeByString(str string) (mediaType int, ok bool) {
	data := map[string]int{
		"picture": PICTURE,
		"video":   VIDEO,
	}
	mediaType, ok = data[str]
	return
}

// TableName sets the insert table name for this struct type
func (p *Media) TableName() string {
	return "medias"
}

func (p *Media) BeforeSave() error {
	return nil
}

func (p *Media) Prepare() {
}

func (p *Media) Validate(action Action) error {

	return nil
}

func AddMedia(mediaURL, userID string, mediaType int) (media Media, err error) {
	db := global.DB
	media.MediaID = uuid.NewV4().String()
	media.UserID = userID
	media.MediaType = mediaType
	media.MediaURL = mediaURL
	err = db.Create(&media).Error
	return
}
