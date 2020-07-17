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
	err = c.BindJSON(&form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "CreateCommunity1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "CreateCommunity2")
		response.FailWithMsg(err.Error(), c)
		return
	}
	userID := c.GetHeader("user_id")
	communityStruct := communityService.CommunityStruct{
		CommunityCreator:     userID,
		CommunityName:        form.CommunityName,
		CommunityDescription: form.CommunityName,
		Background:           form.Background,
	}
	if communityStruct.Exist() {
		response.FailWithMsg(e.REPEAT_SUBMIT, c)
		return
	}
	community, err = communityStruct.Add()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "CreateCommunity3")
		response.FailWithMsg(err.Error(), c)
		return
	}
	data := make(gin.H)
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
	ok, err = model.IsManagerByUserID(userID)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	if !ok {
		response.FailWithMsg(e.PERMISSION_DENY, c)
		return
	}
	err = c.BindJSON(&form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "CreateCommunity1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "CreateCommunity2")
		response.FailWithMsg(err.Error(), c)
		return
	}

	communityStruct := communityService.CommunityStruct{
		CommunityID:          form.CommunityID,
		CommunityName:        form.CommunityName,
		CommunityDescription: form.CommunityDescription,
		Background:           form.Background,
	}
	ok, err = communityStruct.ExistByID()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "CreateCommunity3")
		response.FailWithMsg(err.Error(), c)
		return
	}
	if !ok {
		response.FailWithMsg(e.COMMUNITY_ID_NOT_EXIST, c)
		return
	}
	err = communityStruct.Edit()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "CreateCommunity4")
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
	err = c.BindJSON(&form)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "CollectCommunity1")
		response.FailWithMsg(err.Error(), c)
		return
	}
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
	pageNum, page = utils.GetPage(c)
	userID = c.GetHeader("user_id")
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
