package v1

import (
	"finders-server/global"
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/model/requestForm"
	"finders-server/model/responseForm"
	"finders-server/pkg/e"
	"finders-server/service/collectionService"
	"finders-server/service/communityService"
	"finders-server/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

func CreateCommunity(c *gin.Context) {
	var (
		err       error
		form      requestForm.CreateCommunityForm
		community model.Community
	)
	// 绑定发送过来的json数据
	err = c.BindJSON(&form)
	if err != nil {
		// 若出错则报错
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "CreateCommunity1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	validate := validator.New()
	// 验证数据正确性
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "CreateCommunity2")
		response.FailWithMsg(err.Error(), c)
		return
	}
	// 经过了jwt中间件将user_id 存在了header里 将user_id取出
	userID := c.GetHeader("user_id")
	// 构建结构体 传入数据 通过  service结构体提供的函数来完成操作
	communityStruct := communityService.CommunityStruct{
		CommunityCreator:     userID,
		CommunityName:        form.CommunityName,
		CommunityDescription: form.CommunityName,
		Background:           form.Background,
	}
	// 若重复操作则直接返回
	if communityStruct.Exist() {
		response.FailWithMsg(e.REPEAT_SUBMIT, c)
		return
	}
	// 添加community
	community, err = communityStruct.Add()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "CreateCommunity3")
		response.FailWithMsg(err.Error(), c)
		return
	}
	data := make(gin.H)
	// 返回数据
	data["community_id"] = community.CommunityID
	response.OkWithData(data, c)
}

func UpdateCommunityProfile(c *gin.Context) {
	var (
		err  error
		form requestForm.UpdateCommunityForm
		ok   bool
	)
	userID := c.GetHeader("user_id")
	// 检查是否是manager 若不是manager没有权限修改
	ok, err = model.IsManagerByUserID(userID)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	if !ok {
		response.FailWithMsg(e.PERMISSION_DENY, c)
		return
	}
	// 绑定数据
	err = c.BindJSON(&form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "UpdateCommunityProfile1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "UpdateCommunityProfile2")
		response.FailWithMsg(err.Error(), c)
		return
	}
	// 创建结构体
	communityStruct := communityService.CommunityStruct{
		CommunityID:          form.CommunityID,
		CommunityName:        form.CommunityName,
		CommunityDescription: form.CommunityDescription,
		Background:           form.Background,
	}
	// 检查是否存在 若不存在则错误
	ok, err = communityStruct.ExistByID()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "UpdateCommunityProfile3")
		response.FailWithMsg(err.Error(), c)
		return
	}
	if !ok {
		response.FailWithMsg(e.COMMUNITY_ID_NOT_EXIST, c)
		return
	}
	// 进行修改
	err = communityStruct.Edit()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "UpdateCommunityProfile4")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData("", c)
}

func CollectCommunity(c *gin.Context) {
	var (
		err  error
		form requestForm.GetCommunityIDForm
	)
	// 绑定数据
	err = c.BindJSON(&form)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	validate := validator.New()
	// 检查数据合法性
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "CollectCommunity1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	// 获取中间件设置的user_id
	userID := c.GetHeader("user_id")
	collectionStruct := collectionService.CollectionStruct{
		UserID:         userID,
		Link:           strconv.Itoa(form.CommunityID),
		CollectionType: model.CollectionCommunity,
	}
	if collectionStruct.Exist() {
		response.FailWithMsg(e.REPEAT_SUBMIT, c)
		return
	}
	// 进行收藏
	_, err = collectionStruct.AddCollection()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "CollectCommunity2")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData("", c)
}

func UnCollectCommunity(c *gin.Context) {
	var (
		err  error
		form requestForm.GetCommunityIDForm
	)
	err = c.BindJSON(&form)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "UnCollectCommunity1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	userID := c.GetHeader("user_id")
	collectionStruct := collectionService.CollectionStruct{
		UserID: userID,
		Link:   strconv.Itoa(form.CommunityID),
	}
	if !collectionStruct.Exist() {
		response.FailWithMsg(e.REPEAT_SUBMIT, c)
		return
	}
	err = collectionStruct.RemoveCollection()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "UnCollectCommunity2")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData("", c)
}

func GetCollectCommunity(c *gin.Context) {
	var (
		err     error
		userID  string
		page    int
		pageNum int
		form    responseForm.CommunitiesResponseForm
	)
	// 获取页数 方便分页
	pageNum, page = utils.GetPage(c)
	userID = c.GetHeader("user_id")
	tmpID := c.Query("user_id")
	if tmpID != "" {
		userID = tmpID
	}
	collectionStruct := collectionService.CollectionStruct{
		CollectionType: model.CollectionCommunity,
		UserID:         userID,
		PageNum:        pageNum,
		PageSize:       global.CONFIG.AppSetting.PageSize,
		Page:           page,
	}
	form, err = collectionStruct.GetCommunitiesCollectionResponse()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "GetCollectCommunity1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData(form, c)
}
