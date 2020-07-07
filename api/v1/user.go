package v1

import (
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/pkg/e"
	"finders-server/service/userService"
	"finders-server/utils/reg"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

/*
用户相关接口
*/

func SendCode(c *gin.Context) {
	phone := c.Param("phone")
	if !reg.Phone(phone) {
		response.FailWithMsg("phone not valid", c)
		return
	}
	code, err := userService.SendCode(phone)

	if err != nil {
		response.FailWithMsg("get code fail", c)
		return
	}
	data := gin.H{}
	data["code"] = code
	response.OkWithData(data, c)
}

func Register(user *model.User) (err error) {
	err = userService.RegisterByPhone(user)
	return
}

func Login(c *gin.Context) {
	var (
		user  model.User
		err   error
		token string
	)
	// 假设短信验证是另一个接口
	user, err = userService.CheckByUserNameOrPhone(c)
	// 信息不完全 电话号码和用户名都没有 或信息格式出现错误
	if err != nil {
		// 手机号不存在
		if err.Error() == e.PHONE_NOT_EXIST {
			// 若不存在则自动进行注册
			err = Register(&user)
			if err != nil {
				response.FailWithMsg(err.Error(), c)
				return
			}
		} else {
			response.FailWithMsg(err.Error(), c)
			return
		}
	}
	token, err = userService.GetAuth(user)
	// 获得token失败
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OKWithToken(token, c)
}

func UpdateProfile(c *gin.Context) {
	var (
		user model.User
		err  error
		form userService.UpdateForm
	)
	// 从header获取token
	userName := c.GetHeader("username")
	user, err = model.GetUserByUserName(userName)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	// 获取body中的更新数据
	form, err = userService.GetUpdateForm(c)
	if err != nil {
		//response.FailWithMsg("get update form fail or value not valid", c)
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}

	err = userService.UpdateUserInfo(user, form)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData("", c)
}

func Follow(c *gin.Context) {
	var (
		err              error
		fromUser, toUser model.User
	)
	userName := c.GetHeader("username")
	fromUser, err = model.GetUserByUserName(userName)
	// 解析token错误
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	toUser, err = userService.GetToUser(c)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	_, err = userService.AddRelation(fromUser.UserID, toUser.UserID, model.FOLLOW)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData("", c)
}

func UnFollow(c *gin.Context) {
	var (
		err              error
		fromUser, toUser model.User
	)
	userName := c.GetHeader("username")
	fromUser, err = model.GetUserByUserName(userName)
	// 解析token错误
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	toUser, err = userService.GetToUser(c)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	_, err = userService.DeleteRelation(fromUser.UserID, toUser.UserID, model.FOLLOW)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData("", c)
}

func AddDenyList(c *gin.Context) {
	var (
		err              error
		fromUser, toUser model.User
	)

	userName := c.GetHeader("username")
	fromUser, err = model.GetUserByUserName(userName)
	// 解析token错误
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	toUser, err = userService.GetToUser(c)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	_, err = userService.AddRelation(fromUser.UserID, toUser.UserID, model.DENY)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData("", c)
}

func RemoveDenyList(c *gin.Context) {
	var (
		err              error
		fromUser, toUser model.User
	)
	userName := c.GetHeader("username")
	fromUser, err = model.GetUserByUserName(userName)
	// 解析token错误
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	toUser, err = userService.GetToUser(c)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	_, err = userService.DeleteRelation(fromUser.UserID, toUser.UserID, model.DENY)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	response.OkWithData("", c)
}

func GetDenyList(c *gin.Context) {
	var (
		err             error
		fromUser        model.User
		simpleUserInfos []userService.SimpleUserInfo
	)
	userName := c.GetHeader("username")
	fromUser, err = model.GetUserByUserName(userName)
	// 解析token错误
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	simpleUserInfos, err = userService.GetSimpleUserInfoListByUserID(fromUser.UserID, model.DENY, userService.FROM)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData(simpleUserInfos, c)
}

func GetFans(c *gin.Context) {
	var (
		err             error
		userID          string
		userUUID        uuid.UUID
		simpleUserInfos []userService.SimpleUserInfo
	)
	userID = c.Query("userId")
	if userID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	userUUID, err = uuid.FromString(userID)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	simpleUserInfos, err = userService.GetSimpleUserInfoListByUserID(userUUID, model.FOLLOW, userService.TO)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData(simpleUserInfos, c)
}

func GetFollow(c *gin.Context) {
	var (
		err             error
		userID          string
		userUUID        uuid.UUID
		simpleUserInfos []userService.SimpleUserInfo
	)
	userID = c.Query("userId")
	if userID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	userUUID, err = uuid.FromString(userID)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	simpleUserInfos, err = userService.GetSimpleUserInfoListByUserID(userUUID, model.FOLLOW, userService.FROM)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData(simpleUserInfos, c)
}
