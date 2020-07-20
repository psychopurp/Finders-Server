package requestForm

import (
	"finders-server/global/response"
	"finders-server/pkg/e"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type LoginByUserNameOrPhone struct {
	UserName string `json:"userName" validate:"omitempty,gte=5,lte=30"`
	Password string `json:"password" validate:"omitempty,gte=5,lte=30"`
	Phone    string `json:"phone" validate:"omitempty,gte=1,lte=30"`
	Code     string `json:"code" validate:"omitempty,gte=4,lte=30"`
}

func (u *LoginByUserNameOrPhone) Check(c *gin.Context) bool {
	validate := validator.New()
	err := validate.Struct(*u)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return true
	}
	return false
}

type UpdateUserForm struct {
	Password string `json:"password" validate:"omitempty,gte=1,lte=100"` //[ 2] password                                       VARCHAR[30]          null: false  primary: false  auto: false
	Nickname string `json:"nickname" validate:"omitempty,gte=1,lte=30"`  //[ 3] nickname                                       VARCHAR[30]          null: false  primary: false  auto: false
	UserName string `json:"username" validate:"omitempty,gte=1,lte=50"`
	Avatar   string `json:"avatar" validate:"omitempty,gte=1,lte=200"`
}

type UpdateUserInfoForm struct {
	TrueName      string `json:"truename" validate:"omitempty,gte=1,lte=40"`      //[ 1] truename                                       VARCHAR[40]          strue   primary: false  auto: false
	Address       string `json:"address" validate:"omitempty,gte=1,lte=200"`      //[ 2] address                                        VARCHAR[200]         strue   primary: false  auto: false
	Sex           string `json:"sex" validate:"omitempty,gte=1,lte=4"`            //[ 3] sex                                            VARCHAR[4]           strue   primary: false  auto: false
	Sexual        string `json:"sexual" validate:"omitempty,gte=1,lte=8"`         //[ 4] sexual                                         VARCHAR[8]           strue   primary: false  auto: false
	Feeling       string `json:"feeling" validate:"omitempty,gte=1,lte=20"`       //[ 5] feeling                                        VARCHAR[20]          strue   primary: false  auto: false
	Birthday      string `json:"birthday" validate:"omitempty,gte=1,lte=20"`      //[ 6] birthday                                       VARCHAR[20]          strue   primary: false  auto: false
	Introduction  string `json:"introduction" validate:"omitempty,gte=1,lte=400"` //[ 7] introduction                                   VARCHAR[400]         strue   primary: false  auto: false
	BloodType     string `json:"blood_type" validate:"omitempty,gte=1,lte=8"`     //[ 8] blood_type                                     VARCHAR[8]           strue   primary: false  auto: false
	Eamil         string `json:"eamil" validate:"omitempty,gte=1,lte=60"`         //[ 9] eamil                                          VARCHAR[60]          strue   primary: false  auto: false
	QQ            string `json:"qq" validate:"omitempty,gte=1,lte=30"`            //[10] qq                                             VARCHAR[30]          strue   primary: false  auto: false
	Wechat        string `json:"wechat" validate:"omitempty,gte=1,lte=30"`        //[11] wechat                                         VARCHAR[30]          strue   primary: false  auto: false
	Profession    string `json:"profession" validate:"omitempty,gte=1,lte=60"`    //[12] profession                                     VARCHAR[60]          strue   primary: false  auto: false
	School        string `json:"school" validate:"omitempty,gte=1,lte=30"`        //[13] school                                         VARCHAR[30]          strue   primary: false  auto: false
	Constellation string `json:"constellation" validate:"omitempty,gte=1,lte=40"` //[14] constellation                                  VARCHAR[40]          strue   primary: false  auto: false
	Credit        int    `json:"credit" validate:"omitempty"`                     //[17] credit                                         INT                  sfalse  primary: false  auto: false
	UserTag       string `json:"user_tag" validate:"omitempty,gte=1,lte=65535"`   //[18] user_tag                                       TEXT[65535]          strue   primary: false  auto: false
	Age           int    `json:"age" validate:"omitempty,gte=1"`
}

type UserUpdateForm struct {
	UpdateUserForm
	UpdateUserInfoForm
}

func (u *UserUpdateForm) Check(c *gin.Context) bool {
	validate := validator.New()
	err := validate.Struct(*u)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return true
	}
	return false
}

type ToUserForm struct {
	UserID string `json:"userID" validate:"omitempty,gte=1,lte=50"`
}

func (f *ToUserForm) Check(c *gin.Context) bool {
	validate := validator.New()
	err := validate.Struct(*f)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return true
	}
	return false
}
