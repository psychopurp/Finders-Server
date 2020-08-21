package v1

import (
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/model/requestForm"
	"finders-server/model/responseForm"
	"finders-server/pkg/e"
	"finders-server/service/userService"
	"finders-server/utils"
	"finders-server/utils/reg"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

/*
用户相关接口
*/

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
	userStruct := userService.UserStruct{
		Phone:    form.Phone,
		UserName: form.UserName,
		Password: form.Password,
	}

	if form.UserName != "" {
		ok = true
		// 查看用户名和密码是否正确
		_, exist = userStruct.ExistUserByUserNameAndPassword()
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
		_, exist = userStruct.ExistUserByPhone()
		if !exist {
			_, err = userStruct.Register()
			if utils.FailOnError(e.MYSQL_ERROR, err, c) {
				return
			}
		}
	}
	if !ok {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	token, err = userStruct.GetAuth()
	if utils.FailOnError("", err, c) {
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
		err    error
		form   requestForm.UserUpdateForm
		userID string
	)
	err = c.BindJSON(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	if form.Check(c) {
		return
	}
	userID = c.GetHeader("user_id")
	userStruct := userService.UserStruct{
		UserID: uuid.FromStringOrNil(userID),
	}
	err = userStruct.BindUpdateForm(form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	err = userStruct.Edit()
	if utils.FailOnError(e.MYSQL_ERROR, err, c) {
		return
	}
	response.OkWithData("", c)
}

// @Summary 关注
// @Description 关注
// @Tags 关注
// @Accept json
// @Produce json
// @Param data body userService.ToUserForm true "需要关注的人的ID"
// @Success 200 {string} string "success: {"code": 0, data:"", "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/follow [post]
func Follow(c *gin.Context) {
	var (
		err    error
		userID string
		form   requestForm.ToUserForm
	)
	userID = c.GetHeader("user_id")
	err = c.BindJSON(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	if form.Check(c) {
		return
	}
	userStruct := userService.UserStruct{
		UserID: uuid.FromStringOrNil(userID),
	}
	_, err = userStruct.AddRelation(form.UserID, model.FOLLOW)
	if utils.FailOnError("", err, c) {
		return
	}
	response.OkWithData("", c)
}

// @Summary 取消关注
// @Description 取消关注
// @Tags 取消关注
// @Accept json
// @Produce json
// @Param data body userService.ToUserForm true "需要取消关注的人的ID"
// @Success 200 {string} string "success: {"code": 0, data:"", "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/unfollow [post]
func UnFollow(c *gin.Context) {
	var (
		err    error
		userID string
		form   requestForm.ToUserForm
	)
	userID = c.GetHeader("user_id")
	err = c.BindJSON(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	if form.Check(c) {
		return
	}
	userStruct := userService.UserStruct{
		UserID: uuid.FromStringOrNil(userID),
	}
	_, err = userStruct.DeleteRelation(form.UserID, model.FOLLOW)
	if utils.FailOnError("", err, c) {
		return
	}
	response.OkWithData("", c)
}

// @Summary 拉入黑名单
// @Description 拉入黑名单
// @Tags 拉入黑名单
// @Accept json
// @Produce json
// @Param data body userService.ToUserForm true "需要拉入黑名单的人的ID"
// @Success 200 {string} string "success: {"code": 0, data:"", "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/add_denylist [post]
func AddDenyList(c *gin.Context) {
	var (
		err    error
		userID string
		form   requestForm.ToUserForm
	)
	userID = c.GetHeader("user_id")
	err = c.BindJSON(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	if form.Check(c) {
		return
	}
	userStruct := userService.UserStruct{
		UserID: uuid.FromStringOrNil(userID),
	}
	_, err = userStruct.AddRelation(form.UserID, model.DENY)
	if utils.FailOnError("", err, c) {
		return
	}
	response.OkWithData("", c)
}

// @Summary 将一个人从黑名单中移除
// @Description 将一个人从黑名单中移除
// @Tags 将一个人从黑名单中移除
// @Accept json
// @Produce json
// @Param data body userService.ToUserForm true "需要取消黑名单的人的ID"
// @Success 200 {string} string "success: {"code": 0, data:"", "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/remove_denylist [post]
func RemoveDenyList(c *gin.Context) {
	var (
		err    error
		userID string
		form   requestForm.ToUserForm
	)
	userID = c.GetHeader("user_id")
	err = c.BindJSON(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	userStruct := userService.UserStruct{
		UserID: uuid.FromStringOrNil(userID),
	}
	_, err = userStruct.DeleteRelation(form.UserID, model.DENY)
	if utils.FailOnError("", err, c) {
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
		userID          string
		simpleUserInfos []responseForm.SimpleUserInfo
	)
	userID = c.GetHeader("user_id")
	userStruct := userService.UserStruct{
		UserID: uuid.FromStringOrNil(userID),
	}

	//simpleUserInfos, err = userService.GetSimpleUserInfoListByUserID(fromUser.UserID, model.DENY, userService.FROM)
	simpleUserInfos, err = userStruct.GetSimpleUserInfoListByUserID(model.DENY, userService.FROM)
	if utils.FailOnError("", err, c) {
		return
	}
	response.OkWithData(simpleUserInfos, c)
}

func GetFollowList(c *gin.Context) {
	var (
		err             error
		userID          string
		simpleUserInfos []responseForm.SimpleUserInfo
	)
	userID = c.GetHeader("user_id")
	userStruct := userService.UserStruct{
		UserID: uuid.FromStringOrNil(userID),
	}

	//simpleUserInfos, err = userService.GetSimpleUserInfoListByUserID(fromUser.UserID, model.DENY, userService.FROM)
	simpleUserInfos, err = userStruct.GetSimpleUserInfoListByUserID(model.FOLLOW, userService.FROM)
	if utils.FailOnError("", err, c) {
		return
	}
	response.OkWithData(simpleUserInfos, c)
}

// @Summary 获取一个用户的粉丝列表
// @Description 获取一个用户的粉丝列表
// @Tags 获取一个用户的粉丝列表
// @Accept json
// @Produce json
// @Param data body userService.ToUserForm true "需要获取粉丝列表的人的ID"
// @Success 200 {string} string "success: {"code": 0, data: {"code": 0, data: {userId:"",avatar:"url",nickName:"",introduction:""}, "msg": ""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/get_fans [get]
func GetFans(c *gin.Context) {
	var (
		err             error
		userID          string
		userUUID        uuid.UUID
		simpleUserInfos []responseForm.SimpleUserInfo
	)
	userID = c.Query("userId")
	if userID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	userUUID = uuid.FromStringOrNil(userID)
	userStruct := userService.UserStruct{
		UserID: userUUID,
	}
	//simpleUserInfos, err = userService.GetSimpleUserInfoListByUserID(userUUID, model.FOLLOW, userService.TO)
	simpleUserInfos, err = userStruct.GetSimpleUserInfoListByUserID(model.FOLLOW, userService.TO)
	if utils.FailOnError("", err, c) {
		return
	}
	response.OkWithData(simpleUserInfos, c)
}

// @Summary 获取一个用户的关注列表
// @Description 获取一个用户的关注列表
// @Tags 获取一个用户的关注列表
// @Accept json
// @Produce json
// @Param data body userService.ToUserForm true "需要获取关注列表名单的人的ID"
// @Success 200 {string} string "success: {"code": 0, data: {"code": 0, data: {userId:"",avatar:"url",nickName:"",introduction:""}; failure: {"code": -1, data:"", "msg": "error msg"}"
// @Router /v1/user/get_follow [get]
func GetFollow(c *gin.Context) {
	var (
		err             error
		userID          string
		userUUID        uuid.UUID
		simpleUserInfos []responseForm.SimpleUserInfo
	)
	userID = c.Query("userId")
	if userID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	userUUID = uuid.FromStringOrNil(userID)
	userStruct := userService.UserStruct{
		UserID: userUUID,
	}
	//simpleUserInfos, err = userService.GetSimpleUserInfoListByUserID(userUUID, model.FOLLOW, userService.FROM)
	simpleUserInfos, err = userStruct.GetSimpleUserInfoListByUserID(model.FOLLOW, userService.FROM)
	if utils.FailOnError("", err, c) {
		return
	}
	response.OkWithData(simpleUserInfos, c)
}

func CheckFollow(c *gin.Context) {
	var (
		myID   string
		userID string
		myUUID uuid.UUID
	)
	myID = c.GetHeader("user_id")
	userID = c.Query("userId")
	if userID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	myUUID = uuid.FromStringOrNil(myID)
	userStruct := userService.UserStruct{
		UserID: myUUID,
	}
	flag := userStruct.CheckExistRelation(userID, model.FOLLOW)
	data := make(gin.H)
	data["flag"] = flag
	response.OkWithData(data, c)
}
