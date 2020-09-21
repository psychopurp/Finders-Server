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

type Moment struct {
	MomentID     string `gorm:"column:moment_id;type:varchar(50);primary_key" json:"moment_id"` //[ 0] activity_id                                    VARCHAR[30]          null: false  primary: true   auto: false
	MomentStatus int    `gorm:"column:moment_status;type:INT;" json:"moment_status"`            //[ 1] activity_status                                INT                  null: true   primary: false  auto: false
	MomentInfo   string `gorm:"column:moment_info;type:TEXT;size:65535;" json:"moment_info"`    //[ 2] activity_info                                  TEXT[65535]          null: true   primary: false  auto: false
	// 曝光书
	ReadNum int `gorm:"column:read_num;type:INT;" json:"read_num"` //[ 6] read_num                                       INT                  null: false  primary: false  auto: false
	//ActivityTag    string `gorm:"column:activity_tag;type:TEXT;size:65535;" json:"activity_tag"`          //[ 7] activity_tag                                   TEXT[65535]          null: true   primary: false  auto: false
	MediaIDs string `gorm:"column:media_ids;varchar(5000);" json:"media_ids"` //[ 8] picture_id                                     VARCHAR[30]          null: true   primary: false  auto: false
	Location string `gorm:"column:location;varchar(500);" json:"location"`    //[ 8] picture_id                                     VARCHAR[30]          null: true   primary: false  auto: false
	//MediaType   int    `gorm:"column:media_type;int;" json:"media_id"`       //[ 8] picture_id                                     VARCHAR[30]          null: true   primary: false  auto: false
	User   User   `gorm:"foreignkey:user_id;association_foreignkey:user_id"`
	UserID string `gorm:"column:user_id;type:varchar(50);" json:"user_id"` //[ 9] user_id                                        VARCHAR[30]          null: true   primary: false  auto: false
	TimeModel
}

// TableName sets the insert table name for this struct type
func (a *Moment) TableName() string {
	return "moments"
}

func (a *Moment) BeforeSave() error {
	return nil
}

func (a *Moment) Prepare() {
}

func (a *Moment) Validate(action Action) error {

	return nil
}

// status
const (
	MomentNormal = baseIndex + iota
)

func (a *AffairService) AddMomentByMap(data map[string]interface{}) (moment Moment, err error) {
	db := a.tx
	moment = Moment{
		MomentID:     uuid.NewV4().String(),
		MomentStatus: MomentNormal,
		MomentInfo:   data["moment_info"].(string),
		ReadNum:      0,
		MediaIDs:     data["media_ids"].(string),
		Location:     data["location"].(string),
		UserID:       data["user_id"].(string),
	}
	err = db.Create(&moment).Error
	return
}

func GetMomentsTotal(userID string) (cnt int, err error) {
	db := global.DB
	err = db.Model(&Moment{}).Where("user_id = ?", userID).Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	return
}

func GetMoments(pageNum, pageSize int, userID string) (moments []*Moment, err error) {
	db := global.DB
	err = db.Where("user_id = ?", userID).Offset(pageNum).Limit(pageSize).Find(&moments).Error
	return
}
func GetMomentByMomentID(momentID string) (moment Moment, err error) {
	db := global.DB
	err = db.Where("moment_id = ?", momentID).First(&moment).Error
	return
}

func (a *AffairService) AddMomentReadNum(momentID string) (err error) {
	db := a.tx
	var moment Moment
	err = db.Model(&Moment{}).Where("moment_id = ?", momentID).First(&moment).Error
	if err != nil {
		return
	}
	moment.ReadNum += 1
	err = db.Save(&moment).Error
	return
}
