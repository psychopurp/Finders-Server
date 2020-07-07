package v1

import (
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/pkg/e"
	"finders-server/service/admin_service"
	"github.com/gin-gonic/gin"
)

func AdminRegister(admin *model.Admin) (err error) {
	err = admin_service.RegisterByPhone(admin)
	return
}

func AdminLogin(c *gin.Context) {
	var (
		admin model.Admin
		err   error
		token string
	)

	// 假设短信验证是另一个接口
	admin, err = admin_service.CheckByAdminNameOrPhone(c)
	// 信息不完全 电话号码和用户名都没有 或信息格式出现错误
	if err != nil {
		// 手机号不存在
		if err.Error() == e.PHONE_NOT_EXIST {
			// 若不存在则自动进行注册
			err = AdminRegister(&admin)
			if err != nil {
				response.FailWithMsg(err.Error(), c)
				return
			}
		} else {
			response.FailWithMsg(err.Error(), c)
			return
		}
	}
	token, err = admin_service.GetAuth(admin)
	// 获得token失败
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OKWithToken(token, c)
}

func UpdateAdminProfile(c *gin.Context) {
	var (
		admin model.Admin
		err   error
		form  admin_service.UpdateForm
	)
	// 从header获取token
	adminName := c.GetHeader("adminName")
	admin, err = model.GetAdminByAdminName(adminName)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	// 获取body中的更新数据
	form, err = admin_service.GetUpdateForm(c)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	// 更新admin信息
	err = admin_service.UpdateAdminProfile(admin, form)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData("", c)
}
