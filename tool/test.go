package tool

import (
	"finders-server/model"
	"finders-server/service"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"os"
	"time"
)

type User struct {
	UserID    string         `gorm:"column:user_id;type:varchar(50);primary_key" json:"user_id"`           //[ 0] user_id                                        VARCHAR[30]          null: false  primary: true   auto: false
	Phone     string         `gorm:"column:phone;type:varchar(30);unique_index:unique_phone" json:"phone"` //[ 1] phone                                          VARCHAR[30]          null: false  primary: false  auto: false
	Password  string         `gorm:"column:password;type:varchar(100);" json:"password"`                   //[ 2] password                                       VARCHAR[30]          null: false  primary: false  auto: false
	Nickname  string         `gorm:"column:nickname;type:varchar(30);" json:"nickname"`                    //[ 3] nickname                                       VARCHAR[30]          null: false  primary: false  auto: false
	Status    int            `gorm:"column:status;type:INT;" json:"status"`                                //[ 5] status                                         INT                  null: false  primary: false  auto: false
	CreatedAt time.Time      `gorm:"column:created_at;type:DATETIME;" json:"created_at"`                   //[ 4] created_at                                     DATETIME             null: false  primary: false  auto: false
	UpdatedAt time.Time      `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"`                   //[16] updated_at                                     DATETIME             strue   primary: false  auto: false
	DeletedAt *time.Time     `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`                   //[ 6] deleted_at                                     DATETIME             null: true   primary: false  auto: false
	Avatar    string         `gorm:"column:avatar;type:varchar(100);" json:"avatar"`                       //[ 7] avatar                                         VARCHAR[100]         null: false  primary: false  auto: false
	UserInfo  model.UserInfo `gorm:"foreignkey:UserId"`                                                    //一对一关系
	UserName  string         `gorm:"column:username;type:varchar(50);unique_index:unique_username" json:"userName"`
	//Relations []Relation `gorm:"many2many:relations;foreignkey:from_uid;association_jointable_foreignkey:relation_id"` //多对多关系
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}

type UserInfo struct {
	UserID        string     `gorm:"column:user_id;type:varchar(50);primary_key;" json:"user_id"` //[ 0] user_id                                        VARCHAR[30]          sfalse  primary: true   auto: false
	TrueName      string     `gorm:"column:truename;type:varchar(40);" json:"truename"`           //[ 1] truename                                       VARCHAR[40]          strue   primary: false  auto: false
	Address       string     `gorm:"column:address;type:varchar(200);" json:"address"`            //[ 2] address                                        VARCHAR[200]         strue   primary: false  auto: false
	Sex           string     `gorm:"column:sex;type:varchar(4);" json:"sex"`                      //[ 3] sex                                            VARCHAR[4]           strue   primary: false  auto: false
	Sexual        string     `gorm:"column:sexual;type:varchar(8);" json:"sexual"`                //[ 4] sexual                                         VARCHAR[8]           strue   primary: false  auto: false
	Feeling       string     `gorm:"column:feeling;type:varchar(20);" json:"feeling"`             //[ 5] feeling                                        VARCHAR[20]          strue   primary: false  auto: false
	Birthday      string     `gorm:"column:birthday;type:varchar(20);" json:"birthday"`           //[ 6] birthday                                       VARCHAR[20]          strue   primary: false  auto: false
	Introduction  string     `gorm:"column:introduction;type:varchar(400);" json:"introduction"`  //[ 7] introduction                                   VARCHAR[400]         strue   primary: false  auto: false
	Signature     string     `gorm:"column:signature;type:varchar(400);" json:"signature"`        //[ 7] introduction                                   VARCHAR[400]         strue   primary: false  auto: false
	BloodType     string     `gorm:"column:blood_type;type:varchar(8);" json:"blood_type"`        //[ 8] blood_type                                     VARCHAR[8]           strue   primary: false  auto: false
	Eamil         string     `gorm:"column:eamil;type:varchar(60);" json:"eamil"`                 //[ 9] eamil                                          VARCHAR[60]          strue   primary: false  auto: false
	QQ            string     `gorm:"column:qq;type:varchar(30);" json:"qq"`                       //[10] qq                                             VARCHAR[30]          strue   primary: false  auto: false
	Wechat        string     `gorm:"column:wechat;type:varchar(30);" json:"wechat"`               //[11] wechat                                         VARCHAR[30]          strue   primary: false  auto: false
	Profession    string     `gorm:"column:profession;type:varchar(60);" json:"profession"`       //[12] profession                                     VARCHAR[60]          strue   primary: false  auto: false
	School        string     `gorm:"column:school;type:varchar(30);" json:"school"`               //[13] school                                         VARCHAR[30]          strue   primary: false  auto: false
	Constellation string     `gorm:"column:constellation;type:varchar(40);" json:"constellation"` //[14] constellation                                  VARCHAR[40]          strue   primary: false  auto: false
	CreatedAt     time.Time  `gorm:"column:created_at;type:DATETIME;" json:"created_at"`          //[15] created_at                                     DATETIME             sfalse  primary: false  auto: false
	UpdatedAt     time.Time  `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"`          //[16] updated_at                                     DATETIME             strue   primary: false  auto: false
	Credit        int        `gorm:"column:credit;type:INT;" json:"credit"`                       //[17] credit                                         INT                  sfalse  primary: false  auto: false
	DeletedAt     *time.Time `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`          //[19] deleted_at                                     DATETIME             strue   primary: false  auto: false
	Age           int        `gorm:"column:age;type:INT;" json:"age"`                             //[20] age
}

// TableName sets the insert table name for this struct type
func (u *UserInfo) TableName() string {
	return "user_infos"
}

type Relation struct {
	RelationID    int       `gorm:"AUTO_INCREMENT;column:relation_id;type:INT;primary_key" json:"relation_id"` //[ 0] relation_id                                    INT                  null: false  primary: true   auto: true
	RelationType  int       `gorm:"column:relation_type;type:INT;" json:"relation_type"`                       //[ 1] relation_type                                  INT                  null: false  primary: false  auto: false
	RelationGroup string    `gorm:"column:relation_group;type:varchar(20);" json:"relation_group"`             //[ 2] relation_group                                 VARCHAR[20]          null: false  primary: false  auto: false
	FromUID       string    `gorm:"column:from_uid;type:varchar(50);" json:"from_uid"`                         //[ 3] from_uid                                       VARCHAR[30]          null: false  primary: false  auto: false
	ToUID         string    `gorm:"column:to_uid;type:varchar(50);" json:"to_uid"`                             //[ 4] to_uid                                         VARCHAR[30]          null: false  primary: false  auto: false
	CreatedAt     time.Time `gorm:"column:created_at;type:DATETIME;" json:"created_at"`                        //[ 5] created_at                                     DATETIME             null: false  primary: false  auto: false
	UpdatedAt     time.Time `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"`                        //[ 6] updated_at                                     DATETIME             null: true   primary: false  auto: false
	//DeletedAt     *time.Time `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`                        //[ 6] deleted_at                                     DATETIME             null: true   primary: false  auto: false
}

func (r *Relation) TableName() string {
	return "relations"
}

func Clean() {
	base := service.Base{}
	base.AffairInit(&gin.Context{})
	base.AffairBegin()()
	db := base.Affair.GetTX()
	for i := 1; i <= 14; i++ {
		var (
			err    error
			userID uuid.UUID
			u      string
		)
		userID = uuid.NewV4()
		u = userID.String()
		fmt.Printf("id:%d, uuid:%v\n", i, u)
		var userInfoModel UserInfo
		err = db.Model(&UserInfo{}).Where("user_id = ?", i).First(&userInfoModel).Error
		dealE(err, &base)
		err = db.Unscoped().Delete(&userInfoModel).Error
		dealE(err, &base)
		userInfoModel.UserID = u
		err = db.Save(&userInfoModel).Error
		dealE(err, &base)
		var userModel User
		err = db.Model(&User{}).Where("user_id = ?", i).First(&userModel).Error
		dealE(err, &base)
		err = db.Unscoped().Delete(&userModel).Error
		dealE(err, &base)
		userModel.UserID = u
		err = db.Save(&userModel).Error
		dealE(err, &base)
		var m []*model.CommunityManager
		err = db.Model(&model.CommunityManager{}).Where("manager_id = ?", i).Find(&m).Error
		dealE(err, &base)
		for _, ms := range m {
			ms.ManagerID = u
			err = db.Save(&ms).Error
			dealE(err, &base)
		}
		//err = db.Debug().Model(&model.CommunityManager{}).Where("manager_id IN (?)", []int{i}).Update("manager_id", u).Error
		//dealE(err, &base)
		var mus []*model.CommunityUser
		err = db.Model(&model.CommunityUser{}).Where("user_id = ?", i).Find(&mus).Error
		dealE(err, &base)
		for _, mu := range mus {
			mu.UserID = u
			err = db.Save(&mu).Error
			dealE(err, &base)
		}
		//err = db.Model(&model.CommunityUser{}).Where("user_id = ?", i).Update("user_id", userID.String()).Error
		var cs []*model.Community
		err = db.Model(&model.Community{}).Where("community_creator = ?", i).Find(&cs).Error
		dealE(err, &base)
		for _, c := range cs {
			c.CommunityCreator = u
			err = db.Save(&c).Error
			dealE(err, &base)
		}
		//err = db.Model(&model.Community{}).Where("community_creator = ?", i).Update("community_creator", userID.String()).Error
		//dealE(err, &base)
		var medias []*model.Media
		err = db.Model(&model.Media{}).Where("user_id = ?", i).Find(&medias).Error
		dealE(err, &base)
		for _, c := range medias {
			c.UserID = u
			err = db.Save(&c).Error
			dealE(err, &base)
		}
		//err = db.Model(&model.Media{}).Where("user_id = ?", i).Update("user_id", userID.String()).Error
		//dealE(err, &base)
		var activities []*model.Activity
		err = db.Model(&model.Activity{}).Where("user_id = ?", i).Find(&activities).Error
		dealE(err, &base)
		for _, c := range activities {
			c.UserID = u
			err = db.Save(&c).Error
			dealE(err, &base)
		}
		//err = db.Model(&model.Activity{}).Where("user_id = ?", i).Update("user_id", userID.String()).Error
		//dealE(err, &base)
		var moments []*model.Moment
		err = db.Model(&model.Moment{}).Where("user_id = ?", i).Find(&moments).Error
		dealE(err, &base)
		for _, c := range moments {
			c.UserID = u
			err = db.Save(&c).Error
			dealE(err, &base)
		}
		//err = db.Model(&model.Moment{}).Where("user_id = ?", i).Update("user_id", userID.String()).Error
		//dealE(err, &base)
		var relations []*Relation
		err = db.Model(&Relation{}).Where("from_uid = ?", i).Find(&relations).Error
		dealE(err, &base)
		for _, c := range relations {
			c.FromUID = u
			err = db.Save(&c).Error
			dealE(err, &base)
		}

		err = db.Model(&Relation{}).Where("to_uid = ?", i).Find(&relations).Error
		dealE(err, &base)
		for _, c := range relations {
			c.ToUID = u
			err = db.Save(&c).Error
			dealE(err, &base)
		}
		//err = db.Model(&model.Relation{}).Where("from_uid = ?", i).Update("from_uid", u).Error
		//dealE(err, &base)
		//err = db.Model(&model.Relation{}).Where("to_uid = ?", i).Update("to_uid", u).Error
		//dealE(err, &base)

		var comments []*model.Comment
		err = db.Model(&model.Comment{}).Where("from_uid = ?", i).Find(&comments).Error
		dealE(err, &base)
		for _, c := range comments {
			c.FromUID = u
			err = db.Save(&c).Error
			dealE(err, &base)
		}
		err = db.Model(&model.Comment{}).Where("to_uid = ?", i).Find(&comments).Error
		dealE(err, &base)
		for _, c := range comments {
			c.ToUID = u
			err = db.Save(&c).Error
			dealE(err, &base)
		}
		//err = db.Model(&model.Comment{}).Where("from_uid = ?", i).Update("from_uid", u).Error
		//dealE(err, &base)
		//err = db.Model(&model.Comment{}).Where("to_uid = ?", i).Update("to_uid", u).Error
		//dealE(err, &base)
		var collections []*model.Collection
		err = db.Model(&model.Collection{}).Where("user_id = ?", i).Find(&collections).Error
		dealE(err, &base)
		for _, c := range collections {
			c.UserID = u
			err = db.Save(&c).Error
			dealE(err, &base)
		}
		//err = db.Model(&model.Collection{}).Where("user_id = ?", i).Update("user_id", userID.String()).Error
		//dealE(err, &base)
	}
	base.AffairFinished(&gin.Context{})
}

func dealE(err error, base *service.Base) {
	if err != nil {
		base.AffairRollback()
		fmt.Println(err)
		fmt.Println("--------------------------------------------")
		os.Exit(3)

	}
}
