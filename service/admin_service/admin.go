package admin_service

import (
	"errors"
	"finders-server/model"
	"finders-server/pkg/e"
	"finders-server/utils"
	"finders-server/utils/reg"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"time"
)

type loginByUserNameOrPhone struct {
	UserName string `json:"userName" validate:"omitempty,gte=5,lte=30"`
	Password string `json:"password" validate:"omitempty,gte=5,lte=30"`
	Phone    string `json:"phone" validate:"omitempty,gte=1,lte=30"`
	Code     string `json:"code" validate:"omitempty,gte=4,lte=30"`
}

func CheckByAdminNameOrPhone(c *gin.Context) (admin model.Admin, err error) {
	var json loginByUserNameOrPhone
	var exist bool
	validate := validator.New()
	if c.BindJSON(&json) == nil {
		// 对结构体中的成员进行检测是否符合要求
		err := validate.Struct(json)
		if err != nil {
			return admin, errors.New(e.INFO_ERROR)
		}
		// 若收到的json中用户名不为空
		if json.UserName != "" {
			// 查看用户名和密码是否正确
			admin, exist = model.ExistAdminByAdminNameAndPassword(json.UserName, json.Password)
			// 若存在则返回用户数据
			if exist {
				return admin, nil
			}
			// 若出现错误则是用户名和密码不正确或不存在
			return admin, errors.New(e.USERNAME_NOT_EXIST_OR_PASSWORD_WRONG)
		}
		// 若收到的json中电话号码不为空
		if json.Phone != "" {
			// 检验手机号码正确性
			if !reg.Phone(json.Phone) {
				return admin, errors.New(e.INFO_ERROR)
			}
			// 检测是否存在用户已经使用该手机号注册 假设目前已经通过短信验证
			admin, exist = model.ExistAdminByPhone(json.Phone)
			if exist {
				return admin, nil
			}
			admin.AdminPhone = json.Phone
			return admin, errors.New(e.PHONE_NOT_EXIST)
		}
	}
	// 上面条件都没有满足说明出现了错误
	return admin, errors.New(e.INFO_ERROR)
}

func RegisterByPhone(admin *model.Admin) (err error) {
	//给用户用户ID进行注册
	admin.AdminID = uuid.NewV4().String()
	admin.AdminName = uuid.NewV4().String()[:20]
	admin.AdminPassword = uuid.NewV4().String()[:20]
	admin.Permission = model.SUPER
	err = model.AddAdmin(admin)
	if err != nil {
		return
	}
	return
}

func GetAuth(admin model.Admin) (token string, err error) {
	jwt := utils.NewJWT()
	createAt := time.Now()
	expiredAt := createAt.Add(time.Hour * 5)
	jwtClaims := utils.JWTClaims{
		MapClaims: nil,
		UserName:  admin.AdminName,
		CreatedAt: createAt.Unix(),
		ExpiredAt: expiredAt.Unix(),
	}
	token, err = jwt.GenerateToken(jwtClaims)
	if err != nil {
		return
	}
	model.AddLoginRecord(admin.AdminID, token, createAt, expiredAt)
	return
}

type UpdateForm struct {
	AdminName     string `json:"admin_name" validate:"omitempty,gte=1,lte=30"`
	AdminPassword string `json:"admin_password" validate:"omitempty,gte=1,lte=30"`
	Permission    int    `json:"permission" validate:"omitempty"`
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

func UpdateAdminByStruct(admin model.Admin, it interface{}) (err error) {
	typ := reflect.TypeOf(it)
	val := reflect.ValueOf(it)
	kd := val.Kind()
	// 判断是否是结构体
	if kd != reflect.Struct {
		return errors.New(e.TYPE_ERROR)
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
				err = model.UpdateAdminByUserID(admin.AdminID, tagName, tagVal.String())
				if err != nil {
					return
				}
			}
		case reflect.Int:
			if tagVal.Int() != 0 {
				err = model.UpdateAdminByUserID(admin.AdminID, tagName, tagVal.Int())
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func UpdateAdminProfile(admin model.Admin, form UpdateForm) (err error) {
	err = UpdateAdminByStruct(admin, form)
	if err != nil {
		return errors.New(e.INFO_ERROR)
	}
	return
}
