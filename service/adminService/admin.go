package adminService

import (
	"errors"
	"finders-server/model"
	"finders-server/model/requestForm"
	"finders-server/pkg/e"
	"finders-server/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"time"
)

type AdminStruct struct {
	AdminID       string
	AdminName     string
	AdminPassword string
	AdminPhone    string
	Permission    int
}

func (a *AdminStruct) ExistUserByUserNameAndPassword() (admin model.Admin, exist bool) {
	admin, exist = model.ExistAdminByAdminNameAndPassword(a.AdminName, a.AdminPassword)
	if exist {
		a.AdminID = admin.AdminID
	}
	return
}

func (a *AdminStruct) ExistUserByPhone() (admin model.Admin, exist bool) {
	admin, exist = model.ExistAdminByPhone(a.AdminPhone)
	if exist {
		a.AdminID = admin.AdminID
	}
	return
}

func (a *AdminStruct) Register() (admin model.Admin, err error) {
	data := map[string]interface{}{
		"admin_id":         uuid.NewV4().String(),
		"admin_name":       uuid.NewV4().String(),
		"admin_password":   uuid.NewV4().String()[:20],
		"admin_phone":      a.AdminPhone,
		"admin_permission": model.SUPER,
	}
	admin, err = model.AddAdminByMaps(data)
	if err != nil {
		return
	}
	a.AdminID = admin.AdminID
	return
}

func (a *AdminStruct) GetAuth() (token string, err error) {
	jwt := utils.NewJWT()
	createAt := time.Now()
	expiredAt := createAt.Add(time.Hour * 5)
	jwtClaims := utils.JWTClaims{
		MapClaims: nil,
		UserID:    a.AdminID,
		CreatedAt: createAt.Unix(),
		ExpiredAt: expiredAt.Unix(),
	}
	token, err = jwt.GenerateToken(jwtClaims)
	if err != nil {
		return
	}
	model.AddLoginRecord(a.AdminID, token, createAt, expiredAt)
	return
}

func (a *AdminStruct) BindUpdateForm(form requestForm.UpdateAdminForm) (err error) {
	a.AdminPassword = form.AdminPassword
	a.Permission = form.Permission
	a.AdminName = form.AdminName
	return nil
}

func (a *AdminStruct) Edit() (err error) {
	admin := model.Admin{
		AdminName:     a.AdminName,
		AdminPassword: a.AdminPassword,
		Permission:    a.Permission,
	}
	return model.UpdateAdminByAdmin(a.AdminID, admin)
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
