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

// @Summary 登录或注册
// @Description 登录或注册
// @Tags 登录或注册
// @Accept json
// @Produce json
// @Param data body userService.loginByUserNameOrPhone true "手机号 可选择手机号和验证码登录 或用户名和密码"
// @Success 200 {string} string "success: {"code": 0, data:"", "msg": "", token: "token"}; failure: {"code": -1, data:"", "msg": "error msg", token: ""}"
// @Router /v1/user/login [post]
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

// @Summary 更新用户信息
// @Description 更新用户信息
// @Tags 更新用户信息
// @Accept json
// @Produce json
// @Param data body userService.UpdateForm true "任意字段都可更新，可一次更新多个字段"
// @Success 200 {string} string "success: {"code": 0, data:"", "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/update_profile [post]
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

// @Summary 关注
// @Description 关注
// @Tags 关注
// @Accept json
// @Produce json
// @Param data body userService.FollowForm true "需要关注的人的ID"
// @Success 200 {string} string "success: {"code": 0, data:"", "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/follow [post]
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

// @Summary 取消关注
// @Description 取消关注
// @Tags 取消关注
// @Accept json
// @Produce json
// @Param data body userService.FollowForm true "需要取消关注的人的ID"
// @Success 200 {string} string "success: {"code": 0, data:"", "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/unfollow [post]
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

// @Summary 拉入黑名单
// @Description 拉入黑名单
// @Tags 拉入黑名单
// @Accept json
// @Produce json
// @Param data body userService.FollowForm true "需要拉入黑名单的人的ID"
// @Success 200 {string} string "success: {"code": 0, data:"", "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/add_denylist [post]
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

// @Summary 将一个人从黑名单中移除
// @Description 将一个人从黑名单中移除
// @Tags 将一个人从黑名单中移除
// @Accept json
// @Produce json
// @Param data body userService.FollowForm true "需要取消黑名单的人的ID"
// @Success 200 {string} string "success: {"code": 0, data:"", "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/remove_denylist [post]
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

// @Summary 获取本用户的黑名单列表
// @Description 获取本用户的黑名单列表
// @Tags 获取本用户的黑名单列表
// @Accept json
// @Produce json
// @Success 200 {string} string "success: {"code": 0, data: {userId:"",avatar:"url",nickName:"",introduction:""}, "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/get_denylist [get]
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

// @Summary 获取一个用户的粉丝列表
// @Description 获取一个用户的粉丝列表
// @Tags 获取一个用户的粉丝列表
// @Accept json
// @Produce json
// @Param data body userService.FollowForm true "需要获取粉丝列表的人的ID"
// @Success 200 {string} string "success: {"code": 0, data: {"code": 0, data: {userId:"",avatar:"url",nickName:"",introduction:""}, "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/get_fans [get]
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

// @Summary 获取一个用户的关注列表
// @Description 获取一个用户的关注列表
// @Tags 获取一个用户的关注列表
// @Accept json
// @Produce json
// @Param data body userService.FollowForm true "需要获取关注列表名单的人的ID"
// @Success 200 {string} string "success: {"code": 0, data: {"code": 0, data: {userId:"",avatar:"url",nickName:"",introduction:""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/get_follow [get]
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
