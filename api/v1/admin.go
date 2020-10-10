package v1

import (
	"finders-server/global/response"
	"finders-server/model/requestForm"
	"finders-server/pkg/e"
	"finders-server/service"
	"finders-server/utils"
	"finders-server/utils/reg"
	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	var (
		err   error
		token string
		form  requestForm.LoginByUserNameOrPhone
		exist bool
		ok    bool
	)
	err = c.BindJSON(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	if form.Check(c) {
		return
	}
	adminStruct := service.AdminStruct{
		AdminName:     form.UserName,
		AdminPassword: form.Password,
		AdminPhone:    form.Phone,
	}
	if form.UserName != "" {
		ok = true
		// 查看用户名和密码是否正确
		_, exist = adminStruct.ExistUserByUserNameAndPassword()
		// 若存在则返回用户数据
		if !exist {
			response.FailWithMsg(e.USERNAME_NOT_EXIST_OR_PASSWORD_WRONG, c)
			return
		}
	}
	// 若收到的json中电话号码不为空
	if form.Phone != "" {
		ok = true
		// 检验手机号码正确性
		if !reg.Phone(form.Phone) {
			response.FailWithMsg(e.INFO_ERROR, c)
			return
		}
		// 检测是否存在用户已经使用该手机号注册 假设目前已经通过短信验证
		_, exist = adminStruct.ExistUserByPhone()
		if !exist {
			_, err = adminStruct.Register()
			if utils.FailOnError(e.MYSQL_ERROR, err, c) {
				return
			}
		}
	}
	if !ok {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	token, err = adminStruct.GetAuth()
	if utils.FailOnError("", err, c) {
		return
	}
	response.OKWithToken(token, c)

}

func UpdateAdminProfile(c *gin.Context) {
	var (
		err  error
		form requestForm.UpdateAdminForm
	)
	err = c.BindJSON(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	if form.Check(c) {
		return
	}
	adminID := c.GetHeader("admin_id")
	adminStruct := service.AdminStruct{
		AdminID: adminID,
	}
	err = adminStruct.BindUpdateForm(form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	err = adminStruct.Edit()
	if utils.FailOnError(e.MYSQL_ERROR, err, c) {
		return
	}
	response.OkWithData("", c)
}
