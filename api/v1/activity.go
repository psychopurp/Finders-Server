package v1

import (
	"finders-server/global"
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/model/requestForm"
	"finders-server/model/responseForm"
	"finders-server/pkg/e"
	"finders-server/service/activityService"
	"finders-server/service/collectionService"
	"finders-server/service/commentService"
	"finders-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"gopkg.in/go-playground/validator.v9"
)

func GetActivities(c *gin.Context) {
	var (
		err         error
		communityID int
		pageNum     int
		page        int
		form        responseForm.ActivitiesResponseForm
	)
	communityID = com.StrTo(c.Query("community_id")).MustInt()
	pageNum, page = utils.GetPage(c)
	activityStruct := activityService.ActivityStruct{
		CommunityID: communityID,
		PageNum:     pageNum,
		PageSize:    global.CONFIG.AppSetting.PageSize,
		Page:        page,
	}
	form, err = activityStruct.GetActivitiesPageResponse()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "GetActivities1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData(form, c)
}

func AddActivity(c *gin.Context) {
	var (
		err      error
		form     requestForm.CreateActivityForm
		userID   string
		activity model.Activity
	)
	err = c.BindJSON(&form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "AddActivity1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "AddActivity2")
		response.FailWithMsg(err.Error(), c)
		return
	}
	userID = c.GetHeader("user_id")
	mediaType, ok := model.GetMediaTypeByString(form.MediaType)
	if !ok {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	activityStruct := activityService.ActivityStruct{
		ActivityInfo: form.ActivityInfo,
		MediaID:      form.MediaID,
		MediaType:    mediaType,
		UserID:       userID,
		CommunityID:  form.CommunityID,
	}
	if activityStruct.Exist() {
		response.FailWithMsg(e.REPEAT_SUBMIT, c)
		return
	}
	activity, err = activityStruct.Add()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "AddActivity3")
		response.FailWithMsg(err.Error(), c)
		return
	}
	data := make(gin.H)
	data["activity_id"] = activity.ActivityID
	response.OkWithData(data, c)
}

func GetActivityInfo(c *gin.Context) {
	var (
		err              error
		activityInfoForm responseForm.ActivityInfoForm
		activityID       string
	)
	activityID = c.Query("activity_id")
	if activityID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	activityStruct := activityService.ActivityStruct{
		ActivityID: activityID,
	}
	if !activityStruct.ExistByID() {
		response.FailWithMsg(e.INFO_NOT_EXIST, c)
		return
	}
	err = activityStruct.AddReadNum()
	if err != nil {
		response.FailWithMsg(e.MYSQL_ERROR, c)
		return
	}
	activityInfoForm, err = activityStruct.GetActivityInfoResponse()
	if err != nil {
		response.FailWithMsg(e.MYSQL_ERROR, c)
		return
	}

	response.OkWithData(activityInfoForm, c)
}

func CollectActivity(c *gin.Context) {
	var (
		err    error
		form   requestForm.GetActivityIDForm
		userID string
	)
	err = c.BindJSON(&form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "CollectActivity1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "CollectActivity2")
		response.FailWithMsg(err.Error(), c)
		return
	}
	userID = c.GetHeader("user_id")
	collectionStruct := collectionService.CollectionStruct{
		UserID:         userID,
		Link:           form.ActivityID,
		CollectionType: model.CollectionActivity,
	}
	if collectionStruct.Exist() {
		response.FailWithMsg(e.REPEAT_SUBMIT, c)
		return
	}
	_, err = collectionStruct.AddCollection()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "CollectActivity3")
		response.FailWithMsg(err.Error(), c)
		return
	}
	err = collectionStruct.AddCollectNum()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "CollectActivity4")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData("", c)
}

func UnCollectActivity(c *gin.Context) {
	var (
		err    error
		form   requestForm.GetActivityIDForm
		userID string
	)
	err = c.BindJSON(&form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "UnCollectActivity1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "UnCollectActivity2")
		response.FailWithMsg(err.Error(), c)
		return
	}
	userID = c.GetHeader("user_id")
	collectionStruct := collectionService.CollectionStruct{
		UserID:         userID,
		Link:           form.ActivityID,
		CollectionType: model.CollectionActivity,
	}
	if !collectionStruct.Exist() {
		response.FailWithMsg(e.REPEAT_SUBMIT, c)
		return
	}
	err = collectionStruct.RemoveCollection()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "UnCollectActivity3")
		response.FailWithMsg(err.Error(), c)
		return
	}
	err = collectionStruct.CutCollectNum()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "UnCollectActivity4")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData("", c)
}

func GetCollectActivities(c *gin.Context) {
	var (
		err           error
		userID        string
		pageNum, page int
		form          responseForm.ActivitiesResponseForm
	)
	pageNum, page = utils.GetPage(c)
	userID = c.GetHeader("user_id")
	collectionStruct := collectionService.CollectionStruct{
		CollectionType: model.CollectionActivity,
		UserID:         userID,
		PageNum:        pageNum,
		PageSize:       global.CONFIG.AppSetting.PageSize,
		Page:           page,
	}
	form, err = collectionStruct.GetActivityCollectionResponse()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "GetCollectActivities 1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData(form, c)
}

func GetActivityLike(c *gin.Context) {
	var (
		err           error
		userID        string
		pageNum, page int
		form          responseForm.ActivityLikesResponseForm
	)
	userID = c.GetHeader("user_id")
	pageNum, page = utils.GetPage(c)
	activityStruct := activityService.ActivityStruct{
		UserID:   userID,
		PageNum:  pageNum,
		PageSize: global.CONFIG.AppSetting.PageSize,
		Page:     page,
	}
	form, err = activityStruct.GetActivityLikesResponse()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "GetActivityLike")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData(form, c)
}

func LikeActivity(c *gin.Context) {
	var (
		err    error
		userID string
		form   requestForm.GetActivityIDForm
	)
	err = c.BindJSON(&form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "LikeActivity1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "LikeActivity2")
		response.FailWithMsg(err.Error(), c)
		return
	}
	userID = c.GetHeader("user_id")
	activityStruct := activityService.ActivityStruct{
		ActivityID: form.ActivityID,
		UserID:     userID,
	}
	if activityStruct.ExistLike() {
		response.FailWithMsg(e.REPEAT_SUBMIT, c)
		return
	}
	_, err = activityStruct.Like()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "LikeActivity3")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData("", c)
}

func DisLikeActivity(c *gin.Context) {
	var (
		err    error
		userID string
		form   requestForm.GetActivityIDForm
	)
	err = c.BindJSON(&form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "DisLikeActivity1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "DisLikeActivity2")
		response.FailWithMsg(err.Error(), c)
		return
	}
	userID = c.GetHeader("user_id")
	activityStruct := activityService.ActivityStruct{
		ActivityID: form.ActivityID,
		UserID:     userID,
	}
	if !activityStruct.ExistLike() {
		response.FailWithMsg(e.REPEAT_SUBMIT, c)
		return
	}
	err = activityStruct.DisLike()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "DisLikeActivity3")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData("", c)
}

func AddComment(c *gin.Context) {
	var (
		err     error
		comment model.Comment
		userID  string
		form    requestForm.CreateCommentForm
	)
	err = c.BindJSON(&form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "AddComment1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "AddComment2")
		response.FailWithMsg(err.Error(), c)
		return
	}
	userID = c.GetHeader("user_id")
	commentStruct := commentService.CommentStruct{
		ItemID:   form.ItemID,
		Content:  form.Content,
		FromUID:  userID,
		ItemType: model.CommentOnActivity,
	}
	if commentStruct.Exist() {
		response.FailWithMsg(e.REPEAT_SUBMIT, c)
		return
	}
	comment, err = commentStruct.AddComment()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "AddComment3")
		response.FailWithMsg(err.Error(), c)
		return
	}
	err = commentStruct.AddCommentNum()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "AddComment4")
		response.FailWithMsg(err.Error(), c)
		return
	}
	data := make(gin.H)
	data["comment_id"] = comment.CommentID
	response.OkWithData(data, c)
}

func AddReply(c *gin.Context) {
	var (
		err     error
		comment model.Comment
		userID  string
		form    requestForm.CreateCommentForm
	)
	err = c.BindJSON(&form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "AddReply1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		err = utils.GetErrorAndLog(e.INFO_ERROR, err, "AddReply2")
		response.FailWithMsg(err.Error(), c)
		return
	}
	userID = c.GetHeader("user_id")
	commentStruct := commentService.CommentStruct{
		ItemID:   form.ItemID,
		Content:  form.Content,
		FromUID:  userID,
		ItemType: model.CommentOnComment,
	}
	if commentStruct.Exist() {
		response.FailWithMsg(e.REPEAT_SUBMIT, c)
		return
	}
	comment, err = commentStruct.AddReply()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "AddReply3")
		response.FailWithMsg(err.Error(), c)
		return
	}
	data := make(gin.H)
	data["comment_id"] = comment.CommentID
	response.OkWithData(data, c)
}

func GetActivityComments(c *gin.Context) {
	var (
		err           error
		pageNum, page int
		activityID    string
		form          responseForm.CommentResponseForm
	)
	activityID = c.Query("activity_id")
	if activityID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	pageNum, page = utils.GetPage(c)
	commentStruct := commentService.CommentStruct{
		ItemID:   activityID,
		ItemType: model.CommentOnActivity,
		PageNum:  pageNum,
		PageSize: global.CONFIG.AppSetting.PageSize,
		Page:     page,
	}
	form, err = commentStruct.GetCommentsResponse()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "GetActivityComments1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData(form, c)
}

func GetCommentReplies(c *gin.Context) {
	var (
		err           error
		pageNum, page int
		activityID    string
		form          responseForm.CommentResponseForm
	)
	activityID = c.Query("comment_id")
	if activityID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	pageNum, page = utils.GetPage(c)
	commentStruct := commentService.CommentStruct{
		ItemID:   activityID,
		ItemType: model.CommentOnComment,
		PageNum:  pageNum,
		PageSize: global.CONFIG.AppSetting.PageSize,
		Page:     page,
	}
	form, err = commentStruct.GetCommentsResponse()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "GetCommentReplies1")
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OkWithData(form, c)
}
