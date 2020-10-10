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


CREATE TABLE `communities` (
  `community_id` int(11) NOT NULL COMMENT '社区ID',
  `community_creator` varchar(30) NOT NULL COMMENT '社区创建者（圈主）',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `community_name` varchar(100) NOT NULL COMMENT '社区名称',
  `community_description` text,
  `community_status` int(11) NOT NULL COMMENT '社区状态',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `background` varchar(200) DEFAULT NULL COMMENT '背景',
  `board` text,
  `deleted_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
  PRIMARY KEY (`community_id`),
  KEY `fk_com_user` (`community_creator`),
  CONSTRAINT `fk_com_user` FOREIGN KEY (`community_creator`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Community struct is a row record of the communities table in the employees database
type Community struct {
	CommunityID          int    `gorm:"column:community_id;type:INT;primary_key" json:"community_id"`                    //[ 0] community_id                                   INT                  null: false  primary: true   auto: false
	CommunityCreator     string `gorm:"column:community_creator;varchar(50);" json:"community_creator"`                  //[ 1] community_creator                              VARCHAR[30]          null: false  primary: false  auto: false
	CommunityName        string `gorm:"column:community_name;type:varchar(100);" json:"community_name"`                  //[ 3] community_name                                 VARCHAR[100]         null: false  primary: false  auto: false
	CommunityDescription string `gorm:"column:community_description;type:TEXT;size:65535;" json:"community_description"` //[ 4] community_description                          TEXT[65535]          null: true   primary: false  auto: false
	CommunityStatus      int    `gorm:"column:community_status;type:INT;" json:"community_status"`                       //[ 5] community_status                               INT                  null: false  primary: false  auto: false
	Background           string `gorm:"column:background;type:varchar(200);" json:"background"`                          //[ 7] background                                     VARCHAR[200]         null: true   primary: false  auto: false
	CommunityAvatar      string `gorm:"column:community_avatar;type:varchar(200);" json:"background"`                    //[ 7] background                                     VARCHAR[200]         null: true   primary: false  auto: false
	TimeModel
}

func (c *Community) AfterCreate(scope *gorm.Scope) error {
	manager := new(CommunityManager)
	manager.CommunityID = c.CommunityID
	manager.ManagerID = c.CommunityCreator
	manager.Permission = ManagerNoPermission
	manager.Status = ManagerWaitForCheck

	if err := scope.DB().Create(manager).Error; err != nil {
		return err
	}
	communityUser := new(CommunityUser)
	communityUser.CommunityID = c.CommunityID
	communityUser.Status = CommunityUserWaitForCheck
	communityUser.UserID = c.CommunityCreator
	if err := scope.DB().Create(communityUser).Error; err != nil {
		return err
	}
	return nil
}

// TableName sets the insert table name for this struct type
func (c *Community) TableName() string {
	return "communities"
}

func (c *Community) BeforeSave() error {
	return nil
}

func (c *Community) Prepare() {
}

func (c *Community) Validate(action Action) error {

	return nil
}

const (
	CommunityWaitForCheck = baseIndex + iota
	CommunityNormal
)

func GetCommunityStatusByString(status string) (int, bool) {
	data := map[string]int{
		"waitForCheck": CommunityUserWaitForCheck,
		"normal":       CommunityNormal,
	}
	value, ok := data[status]
	return value, ok
}

func AddCommunityByMap(data *Community) (community Community, err error) {
	db := global.DB

	data.CommunityStatus = CommunityWaitForCheck
	err = db.Create(&data).Error
	return *data, err
}

func ExistCommunityByMap(data map[string]interface{}) bool {
	db := global.DB
	var community Community
	err := db.Where(data).First(&community).Error
	return !gorm.IsRecordNotFoundError(err)
}

func ExistCommunityByID(id int) (bool, error) {
	db := global.DB
	var community Community
	err := db.Where("community_id = ?", id).First(&community).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return true, nil
}

func UpdateCommunityByCommunity(ID int, community Community) (err error) {
	db := global.DB
	err = db.Model(&Community{}).Where("community_id = ?", ID).Updates(community).Error
	return
}

func GetCommunities(pageNum, pageSize int, communityIDs []string) (communities []*Community, err error) {
	db := global.DB
	err = db.Where("community_id IN (?)", communityIDs).Offset(pageNum).Limit(pageSize).Find(&communities).Error
	return
}

func GetCommunityByCommunityID(communityID int) (community Community, err error) {
	db := global.DB
	err = db.Where("community_id = ?", communityID).First(&community).Error
	return
}

func GetCommunityTotal(communityID int) (cnt int, err error) {
	db := global.DB
	err = db.Model(&Community{}).Where("community_id = ?", communityID).Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	return
}
