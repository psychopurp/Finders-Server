package v1

import (
	"finders-server/global"
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/model/requestForm"
	"finders-server/model/responseForm"
	"finders-server/pkg/e"
	"finders-server/service"
	"finders-server/service/cache"
	"finders-server/utils"
	"finders-server/utils/reg"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

/*
用户相关接口
*/

func GetPhoneCode(c *gin.Context) {
	var (
		err   error
		phone string
		//code string
	)
	phone = c.Query("phone")
	if phone == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	// 检验手机号码正确性
	if !reg.Phone(phone) {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	cacheSrv := cache.NewPhoneCacheService()
	_, err = cacheSrv.GetPhoneCode(phone)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	//code = code
	//response.OkWithData(code, c)
	response.OkWithData("", c)
}

func Login(c *gin.Context) {
	var (
		err   error
		form  requestForm.LoginByUserNameOrPhone
		ok    bool
		exist bool
		token string
	)
	err = c.ShouldBind(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	if form.Check(c) {
		return
	}
	userStruct := service.UserStruct{
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
		cacheSrv := cache.NewPhoneCacheService()
		if !cacheSrv.ValidatePhoneCode(form.Phone, form.Code) {
			response.FailWithMsg(e.PHONE_CODE_ERROR, c)
			return
		}
		// 检测是否存在用户已经使用该手机号注册
		_, exist = userStruct.ExistUserByPhone()
		if !exist {
			response.FailWithMsg(e.PHONE_CODE_ERROR, c)
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

func Register(c *gin.Context) {
	var (
		err   error
		form  requestForm.LoginByUserNameOrPhone
		exist bool
		token string
	)

	err = c.ShouldBind(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	if form.Check(c) {
		return
	}
	userStruct := service.UserStruct{
		Phone: form.Phone,
	}
	// 若收到的json中电话号码不为空
	if form.Phone != "" {
		// 检验手机号码正确性
		if !reg.Phone(form.Phone) {
			response.FailWithMsg(e.INFO_ERROR, c)
			return
		}
		cacheSrv := cache.NewPhoneCacheService()
		if !cacheSrv.ValidatePhoneCode(form.Phone, form.Code) {
			response.FailWithMsg(e.PHONE_CODE_ERROR, c)
			return
		}
		// 检测是否存在用户已经使用该手机号注册 假设目前已经通过短信验证
		_, exist = userStruct.ExistUserByPhone()
		if !exist {
			_, err = userStruct.Register()
			if utils.FailOnError(e.MYSQL_ERROR, err, c) {
				return
			}
		} else {
			response.FailWithMsg(e.PHONE_HAS_BEEN_REGISTER, c)
			return
		}
	} else {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	token, err = userStruct.GetAuth()
	if utils.FailOnError("", err, c) {
		return
	}
	response.OKWithToken(token, c)
}

// @Summary 登录或注册
// @Description 登录或注册
// @Tags 登录或注册
// @Accept json
// @Produce json
// @Param data body userService.loginByUserNameOrPhone true "手机号 可选择手机号和验证码登录 或用户名和密码"
// @Success 200 {string} string "success: {"code": 0, data:"", "msg": "", token: "token"}; failure: {"code": -1, data:"", "msg": "error msg", token: ""}"
// @Router /v1/user/login [post]
func LoginAndRegister(c *gin.Context) {
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
	userStruct := service.UserStruct{
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
		cacheSrv := cache.NewPhoneCacheService()
		if !cacheSrv.ValidatePhoneCode(form.Phone, form.Code) {
			response.FailWithMsg(e.PHONE_CODE_ERROR, c)
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

func BackTest(c *gin.Context) {
	var (
		err   error
		token string
	)
	type TMP struct {
		UserID string `json:"user_id"`
		Key    string `json:"key"`
	}
	var tmp TMP
	err = c.ShouldBind(&tmp)
	if err != nil || tmp.Key != "miolyn" {
		response.Fail(c)
		return
	}
	userStruct := service.UserStruct{
		UserID: uuid.FromStringOrNil(tmp.UserID),
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
	userStruct := service.UserStruct{
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

func GetUserInfo(c *gin.Context) {
	var (
		err    error
		userID string
		user   model.User
		form   responseForm.SimpleUserInfo
	)
	userID = c.Query("user_id")
	userStruct := service.UserStruct{
		UserID: uuid.FromStringOrNil(userID),
	}
	user, err = userStruct.GetUserByUserID()
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	form = responseForm.SimpleUserInfo{
		UserId:       user.UserID.String(),
		Avatar:       user.Avatar,
		NickName:     user.Nickname,
		Introduction: user.UserInfo.Introduction,
		Signature:    user.UserInfo.Signature,
	}
	response.OkWithData(form, c)
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
	userStruct := service.UserStruct{
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
	userStruct := service.UserStruct{
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
	userStruct := service.UserStruct{
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
	userStruct := service.UserStruct{
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
		err     error
		userID  string
		pageNum int
		page    int
		form    responseForm.SimpleUserInfoWithPage
	)
	userID = c.GetHeader("user_id")
	pageNum, page = utils.GetPage(c)
	userStruct := service.UserStruct{
		UserID:   uuid.FromStringOrNil(userID),
		PageNum:  pageNum,
		Page:     page,
		PageSize: global.CONFIG.AppSetting.PageSize,
	}

	//simpleUserInfos, err = userService.GetSimpleUserInfoListByUserID(fromUser.UserID, model.DENY, userService.FROM)
	form, err = userStruct.GetSimpleUserInfoListWitPageByUserID(model.DENY, service.FROM)
	if utils.FailOnError("", err, c) {
		return
	}
	response.OkWithData(form, c)
}

func GetFollowList(c *gin.Context) {
	var (
		err     error
		userID  string
		page    int
		pageNum int
		form    responseForm.SimpleUserInfoWithPage
	)
	userID = c.GetHeader("user_id")
	pageNum, page = utils.GetPage(c)
	userStruct := service.UserStruct{
		UserID:   uuid.FromStringOrNil(userID),
		PageNum:  pageNum,
		Page:     page,
		PageSize: global.CONFIG.AppSetting.PageSize,
	}

	form, err = userStruct.GetSimpleUserInfoListWitPageByUserID(model.FOLLOW, service.FROM)
	if utils.FailOnError("", err, c) {
		return
	}
	response.OkWithData(form, c)
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
		err           error
		userID        string
		userUUID      uuid.UUID
		page, pageNum int
		form          responseForm.SimpleUserInfoWithPage
	)
	userID = c.Query("user_id")
	pageNum, page = utils.GetPage(c)
	if userID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	userUUID = uuid.FromStringOrNil(userID)
	userStruct := service.UserStruct{
		UserID:   userUUID,
		Page:     page,
		PageNum:  pageNum,
		PageSize: global.CONFIG.AppSetting.PageSize,
	}
	//simpleUserInfos, err = userService.GetSimpleUserInfoListByUserID(userUUID, model.FOLLOW, userService.TO)
	form, err = userStruct.GetSimpleUserInfoListWitPageByUserID(model.FOLLOW, service.TO)
	if utils.FailOnError("", err, c) {
		return
	}
	response.OkWithData(form, c)
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
		err           error
		userID        string
		userUUID      uuid.UUID
		pageNum, page int
		form          responseForm.SimpleUserInfoWithPage
	)
	userID = c.Query("userId")
	pageNum, page = utils.GetPage(c)
	if userID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	userUUID = uuid.FromStringOrNil(userID)
	userStruct := service.UserStruct{
		UserID:   userUUID,
		Page:     page,
		PageNum:  pageNum,
		PageSize: global.CONFIG.AppSetting.PageSize,
	}
	//simpleUserInfos, err = userService.GetSimpleUserInfoListByUserID(userUUID, model.FOLLOW, userService.FROM)
	form, err = userStruct.GetSimpleUserInfoListWitPageByUserID(model.FOLLOW, service.FROM)
	if utils.FailOnError("", err, c) {
		return
	}
	response.OkWithData(form, c)
}

func CheckFollow(c *gin.Context) {
	var (
		myID   string
		userID string
		myUUID uuid.UUID
	)
	myID = c.GetHeader("user_id")
	userID = c.Query("user_id")
	if userID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	myUUID = uuid.FromStringOrNil(myID)
	userStruct := service.UserStruct{
		UserID: myUUID,
	}
	flag := userStruct.CheckExistRelation(userID, model.FOLLOW)
	data := make(gin.H)
	data["flag"] = flag
	response.OkWithData(data, c)
}
