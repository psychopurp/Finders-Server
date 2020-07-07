package userService

import (
	"encoding/json"
	"errors"
	"finders-server/global"
	"finders-server/model"
	"finders-server/pkg/e"
	"finders-server/service/cache_service"
	"finders-server/service/redis"
	"finders-server/utils"
	"finders-server/utils/reg"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"time"
)

/*
用户相关的一些service
*/

func RegisterByPhone(user *model.User) (err error) {
	//给用户用户ID进行注册
	user.UserID = uuid.NewV4()
	user.UserName = uuid.NewV4().String()[:20]
	user.Nickname = uuid.NewV4().String()[:20]
	user.Password = uuid.NewV4().String()[:20]
	user.Status = model.Normal
	err = model.AddUser(user)
	if err != nil {
		return
	}
	return
}

type loginByUserNameOrPhone struct {
	UserName string `json:"userName" validate:"omitempty,gte=5,lte=30"`
	Password string `json:"password" validate:"omitempty,gte=5,lte=30"`
	Phone    string `json:"phone" validate:"omitempty,gte=1,lte=30"`
	Code     string `json:"code" validate:"omitempty,gte=4,lte=30"`
}

func CheckByUserNameOrPhone(c *gin.Context) (user model.User, err error) {
	var json loginByUserNameOrPhone
	validate := validator.New()
	if c.BindJSON(&json) == nil {
		// 对结构体中的成员进行检测是否符合要求
		err := validate.Struct(json)
		if err != nil {
			return user, errors.New(e.INFO_ERROR)
		}
		// 若收到的json中用户名不为空
		if json.UserName != "" {
			// 查看用户名和密码是否正确
			user, exist := model.ExistUserByUserNameAndPassword(json.UserName, json.Password)
			// 若存在则返回用户数据
			if exist {
				return user, nil
			}
			// 若出现错误则是用户名和密码不正确或不存在
			return user, errors.New(e.USERNAME_NOT_EXIST_OR_PASSWORD_WRONG)
		}
		// 若收到的json中电话号码不为空
		if json.Phone != "" {
			// 检验手机号码正确性
			if !reg.Phone(json.Phone) {
				return user, errors.New(e.INFO_ERROR)
			}
			// 检测是否存在用户已经使用该手机号注册 假设目前已经通过短信验证
			user, exist := model.ExistUserByPhone(json.Phone)
			if exist {
				return user, nil
			}
			user.Phone = json.Phone
			return user, errors.New(e.PHONE_NOT_EXIST)
		}
	}
	// 上面条件都没有满足说明出现了错误
	return user, errors.New(e.INFO_ERROR)
}

// 通过用户名获取 对应的token
func GetAuth(user model.User) (token string, err error) {
	jwt := utils.NewJWT()
	createAt := time.Now()
	expiredAt := createAt.Add(time.Hour * 5)
	jwtClaims := utils.JWTClaims{
		MapClaims: nil,
		UserName:  user.UserName,
		CreatedAt: createAt.Unix(),
		ExpiredAt: expiredAt.Unix(),
	}
	token, err = jwt.GenerateToken(jwtClaims)
	if err != nil {
		return
	}
	model.AddLoginRecord(user.UserID.String(), token, createAt, expiredAt)
	return
}

func GetUserFromAuth(token string) (user model.User, err error) {
	jwt := utils.NewJWT()
	// 解析token
	jwtClaims, err := jwt.ParseToken(token)
	if err != nil || jwtClaims == nil {
		return user, errors.New(e.TOKEN_ERROR)
	}
	// 检测是否超时
	if jwtClaims.ExpiredAt < time.Now().Unix() {
		return user, errors.New(e.TOKEN_OUT_OF_DATE)
	}
	// 根据用户名获取user
	user, err = model.GetUserByUserName(jwtClaims.UserName)
	if err != nil {
		return user, errors.New(e.MYSQL_ERROR)
	}
	return user, nil
}

func SendCode(phone string) (string, error) {
	var (
		code string
		err  error
	)
	code = GetCacheCode(phone)
	if code == "" {
		code = utils.GetRandomCode()
	}

	cache := cache_service.Phone{Phone: phone}
	key := cache.GetPhoneCodeKey()

	if err = redis.Set(key, code, 600); err != nil {
		global.LOG.Warning("cache set fail", err.Error())
	}
	return code, err
}

func GetCacheCode(phone string) string {
	cache := cache_service.Phone{Phone: phone}
	key := cache.GetPhoneCodeKey()
	if !redis.Exists(key) {
		return ""
	}
	var code string
	data, err := redis.Get(key)
	if err != nil {
		global.LOG.Warning("cache get fail", err.Error())
		return ""
	}
	json.Unmarshal(data, &code)
	return code
}

type UpdateUserForm struct {
	Password string `json:"password" validate:"omitempty,gte=1,lte=100"` //[ 2] password                                       VARCHAR[30]          null: false  primary: false  auto: false
	Nickname string `json:"nickname" validate:"omitempty,gte=1,lte=30"`  //[ 3] nickname                                       VARCHAR[30]          null: false  primary: false  auto: false
	UserName string `json:"username" validate:"omitempty,gte=1,lte=50"`
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

type UpdateForm struct {
	UpdateUserForm
	UpdateUserInfoForm
}

func Debug(a ...interface{}) {
	global.LOG.Debug(a)
}

func GetUpdateForm(c *gin.Context) (form UpdateForm, err error) {
	err = c.BindJSON(&form)
	if err != nil {
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	return
}

func UpdateUserByStruct(user model.User, it interface{}) (err error) {
	typ := reflect.TypeOf(it)
	val := reflect.ValueOf(it)
	kd := val.Kind()
	// 判断是否是结构体
	if kd != reflect.Struct {
		return errors.New("type error, struct needed")
	}
	// 获取结构体的字段数目
	num := val.NumField()
	for i := 0; i < num; i++ {
		// 获取字段名称
		tagName := typ.Field(i).Tag.Get("json")
		// 字段的值
		tagVal := val.Field(i)
		if tagName == "" {
			continue
		}
		// 根据字段类型更新字段的值
		switch tagVal.Kind() {
		case reflect.String:
			if tagVal.String() != "" {
				// 通过UserID找到user并更新对应字段的值
				err = model.UpdateUserByUserID(user.UserID.String(), tagName, tagVal.String())
				if err != nil {
					return
				}
			}
		case reflect.Int:
			if tagVal.Int() != 0 {
				err = model.UpdateUserByUserID(user.UserID.String(), tagName, tagVal.Int())
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func UpdateUserInfoByStruct(user model.User, it interface{}) (err error) {
	typ := reflect.TypeOf(it)
	val := reflect.ValueOf(it)
	kd := val.Kind()
	// 判断是否为结构体
	if kd != reflect.Struct {
		return errors.New("type error, struct needed")
	}
	num := val.NumField()
	for i := 0; i < num; i++ {
		// 获得字段名称
		tagName := typ.Field(i).Tag.Get("json")
		tagVal := val.Field(i)
		if tagName == "" {
			continue
		}
		switch tagVal.Kind() {
		case reflect.String:
			if tagVal.String() != "" {
				// 通过UserID更新userInfo
				err = model.UpdateUserInfoByUserID(user.UserID.String(), tagName, tagVal.String())
				if err != nil {
					return
				}
			}
		case reflect.Int:
			if tagVal.Int() == 0 {
				err = model.UpdateUserInfoByUserID(user.UserID.String(), tagName, tagVal.Int())
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func UpdateUserInfo(user model.User, form UpdateForm) (err error) {
	err = UpdateUserByStruct(user, form.UpdateUserForm)
	if err != nil {
		return errors.New(e.INFO_ERROR)
	}
	err = UpdateUserInfoByStruct(user, form.UpdateUserInfoForm)
	if err != nil {
		return errors.New(e.INFO_ERROR)
	}
	return
}

type FollowForm struct {
	UserID string `json:"userID" validate:"omitempty,gte=1,lte=50"`
}

func GetToUser(c *gin.Context) (user model.User, err error) {
	var json FollowForm
	err = c.BindJSON(&json)
	if err != nil {
		return
	}
	validate := validator.New()
	err = validate.Struct(json)
	if err != nil {
		return
	}
	user, err = model.GetUserByUserID(json.UserID)
	if err != nil {
		return
	}
	return
}

func AddRelation(fromID, toID uuid.UUID, relationType int) (relation model.Relation, err error) {
	var (
		isExist bool
	)
	// 若想添加 黑名单关系
	if relationType == model.DENY {
		// 看看两者有没有关注关系
		isExist, err = model.ExistRelation(fromID, toID, model.FOLLOW)
		if err != nil {
			return relation, errors.New(e.MYSQL_ERROR)
		}
		// 若有 可以直接修改关注为拉黑
		if isExist {
			data := make(map[string]interface{})
			data["from_uid"] = fromID.String()
			data["to_uid"] = toID.String()
			data["relation_type"] = model.FOLLOW
			relation, err = model.UpdateRelationType(data, model.DENY)
			return
		}
	} else {
		// 若是要添加关注关系
		// 先看是否已存在关注关系
		isExist, err = model.ExistRelation(fromID, toID, relationType)
		if err != nil {
			return relation, errors.New(e.MYSQL_ERROR)
		}
		// 若存在则
		if isExist {
			return
		}
	}
	relation, err = model.AddRelation(fromID, toID, relationType)
	if err != nil {
		return relation, errors.New(e.MYSQL_ERROR)
	}
	return
}

func DeleteRelation(fromID, toID uuid.UUID, relationType int) (relation model.Relation, err error) {
	var (
		isExist bool
	)
	isExist, err = model.ExistRelation(fromID, toID, relationType)
	if err != nil {
		return
	}
	if !isExist {
		return
	}
	relation, err = model.DeleteRelation(fromID, toID, relationType)
	return
}

type SimpleUserInfo struct {
	UserId       uuid.UUID
	Avatar       string
	NickName     string
	Introduction string
}

const (
	FROM = 1
	TO   = 2
)

func GetSimpleUserInfoListByUserID(fromID uuid.UUID, relationType int, loc int) (simpleUserInfo []SimpleUserInfo, err error) {
	var (
		isExist   bool
		relations []model.Relation
		users     []model.User
		userInfo  model.UserInfo
	)
	data := make(map[string]interface{})
	data["relation_type"] = relationType
	// 根据粉丝列表还是关注列表选择 userID是属于from还是to
	if loc == FROM {
		data["from_uid"] = fromID.String()
	} else {
		data["to_uid"] = fromID.String()
	}
	isExist, err = model.ExistRelationByData(data)
	if err != nil {
		return
	}
	if !isExist {
		return
	}
	relations, err = model.GetRelationsByData(data)
	if err != nil {
		return
	}
	var userIDs []string
	for _, relation := range relations {
		if loc == FROM {
			userIDs = append(userIDs, relation.ToUID.String())
		} else {
			userIDs = append(userIDs, relation.FromUID.String())
		}
	}
	users, err = model.GetUsersByUserIDs(userIDs)
	for _, user := range users {
		userInfo, err = model.GetUserInfoByUserID(user.UserID.String())
		if err != nil {
			return simpleUserInfo, errors.New(e.MYSQL_ERROR)
		}
		simpleUserInfo = append(simpleUserInfo, SimpleUserInfo{
			UserId:       user.UserID,
			Avatar:       user.Avatar,
			NickName:     user.Nickname,
			Introduction: userInfo.Introduction,
		})
	}
	return
}
