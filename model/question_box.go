package model

import (
	"database/sql"
	"finders-server/global"
	"github.com/guregu/null"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

//
//type MessageBox struct {
//	MessageBoxID     int `gorm:"column:message_box_id;type:int;primary_key" json:"message_box_id"`
//	UserID           string
//	MessageBoxStatus int    // 封禁或正常
//	MessageBoxInfo   string // 提问箱内容
//	UseNum           int    // 使用次数
//	ReplyNum         int    // 回答次数
//	LikeNum          int    // 喜欢的次数
//	TagNames string
//	TimeModel
//}
type QuestionBox struct {
	QuestionBoxID     int    `gorm:"column:question_box_id" json:"question_box_id" form:"question_box_id"`
	UserID            string `gorm:"column:user_id" json:"user_id" form:"user_id"`
	QuestionBoxStatus int    `gorm:"column:question_box_status" json:"question_box_status" form:"question_box_status"`
	QuestionBoxInfo   string `gorm:"column:question_box_info" json:"question_box_info" form:"question_box_info"`
	UseNum            int    `gorm:"column:use_num" json:"use_num" form:"use_num"`
	ReplyNum          int    `gorm:"column:reply_num" json:"reply_num" form:"reply_num"`
	LikeNum           int    `gorm:"column:like_num" json:"like_num" form:"like_num"`
	TagNames          string `gorm:"column:tag_names" json:"tag_names" form:"tag_names"`
	TimeModel
}

// TableName sets the insert table name for this struct type
func (a *QuestionBox) TableName() string {
	return "question_box"
}

func (a *QuestionBox) BeforeSave() error {
	return nil
}

func (a *QuestionBox) Prepare() {
}

func (a *QuestionBox) Validate(action Action) error {

	return nil
}

// status
const (
	QuestionStatusNormal = baseIndex + iota
	QuestionStatusBan
)

func (a *AffairService) AddQuestionBox(questionBox *QuestionBox) (err error) {
	db := a.tx
	err = db.Model(&QuestionBox{}).Create(questionBox).Error
	return
}

func ExistQuestionBoxByInfo(info string) bool {
	db := global.DB
	var questionBox QuestionBox
	err := db.Model(&QuestionBox{}).Where("question_box_info = ?", info).First(&questionBox).Error
	if gorm.IsRecordNotFoundError(err) {
		return false
	}
	return true
}
