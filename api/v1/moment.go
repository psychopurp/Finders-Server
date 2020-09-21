package v1

import (
	"finders-server/global"
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/model/requestForm"
	"finders-server/model/responseForm"
	"finders-server/pkg/e"
	"finders-server/service/baseService"
	"finders-server/service/momentService"
	"finders-server/utils"
	"github.com/gin-gonic/gin"
)

func CreateMoment(c *gin.Context) {
	var (
		err      error
		form     requestForm.CreateMomentRequestForm
		userID   string
		mediaIDs string
		moment   model.Moment
	)
	userID = c.GetHeader("user_id")
	err = c.ShouldBind(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	if form.Check(c) {
		return
	}
	for _, mediaId := range form.MediaIDs {
		mediaIDs = mediaId + ";"
	}
	if len(mediaIDs) == 0 {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	mediaIDs = mediaIDs[:len(mediaIDs)-1]
	momentStruct := momentService.MomentStruct{
		MomentInfo: form.MomentInfo,
		MediaIDs:   mediaIDs,
		Location:   form.Location,
		UserID:     userID,
	}
	if momentStruct.AffairInit(c) {
		return
	}
	momentStruct.AffairBegin()()
	moment, err = momentStruct.Add()
	if momentStruct.AffairRollbackIfError(err, c) {
		return
	}
	if momentStruct.AffairFinished(c) {
		return
	}
	data := make(gin.H)
	data["moment_id"] = moment.MomentID
	response.OkWithData(data, c)
}

func GetUserMoments(c *gin.Context) {
	var (
		err           error
		form          responseForm.GetUserMomentsResponseForm
		userID        string
		pageNum, page int
	)
	userID = c.Query("userId")
	if userID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	pageNum, page = utils.GetPage(c)
	momentStruct := momentService.MomentStruct{
		UserID: userID,
		Base: baseService.Base{
			PageNum:  pageNum,
			PageSize: global.CONFIG.AppSetting.PageSize,
			Page:     page,
			Affair:   nil,
		},
	}
	base := baseService.Base{}
	if base.AffairInit(c) {
		return
	}
	base.AffairBegin()()
	momentStruct.AffairInitWithAffair(base.Affair)
	form, err = momentStruct.GetUserMomentsResponseForm()
	if base.AffairRollbackIfError(err, c) {
		return
	}
	if base.AffairFinished(c) {
		return
	}
	response.OkWithData(form, c)
}

func GetMoment(c *gin.Context) {
	var (
		err      error
		momentID string
		form     responseForm.GetMomentResponseForm
	)
	momentID = c.Query("moment_id")
	if momentID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	momentStruct := momentService.MomentStruct{
		MomentID: momentID,
	}
	form, err = momentStruct.GetMomentResponseForm()
	if utils.FailOnError(e.MYSQL_ERROR, err, c) {
		return
	}
	response.OkWithData(form, c)
}

func LikeMoment(c *gin.Context) {
	var (
		err    error
		form   requestForm.LikeMomentRequestForm
		userID string
	)
	userID = c.GetHeader("user_id")
	err = c.ShouldBind(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	if form.Check(c) {
		return
	}
	momentStruct := momentService.MomentStruct{
		MomentID: form.MomentID,
		UserID:   userID,
	}
	err = momentStruct.Like()
	if utils.FailOnError(e.MYSQL_ERROR, err, c) {
		return
	}
	response.OkWithData("", c)
}
