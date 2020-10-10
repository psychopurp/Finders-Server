package v1

import (
	"finders-server/global/response"
	"finders-server/model/requestForm"
	"finders-server/pkg/e"
	"finders-server/service"
	"finders-server/utils"
	"github.com/gin-gonic/gin"
)

func CreateQuestionBox(c *gin.Context) {
	var (
		userID   string
		err      error
		form     requestForm.CreateQuestionBoxForm
		tagIDs   []int
		tagNames string
	)
	userID = c.GetHeader("user_id")
	err = c.ShouldBind(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c) {
		return
	}
	if form.Check(c) {
		return
	}
	base := service.Base{}
	if base.AffairInit(c) {
		return
	}
	tagNames = ""
	for _, tagName := range form.TagNames {
		tagNames = tagName + ";" + tagNames
	}
	questionStruct := service.QuestionBoxStruct{
		UserID:          userID,
		QuestionBoxInfo: form.QuestionBoxInfo,
		TagNames:        tagNames,
	}
	if questionStruct.CheckExistQuestionInfo() {
		response.FailWithMsg(e.REPEAT_SUBMIT, c)
		return
	}
	base.AffairBegin()()
	tagStruct := service.TagStruct{}
	tagStruct.AffairInitWithAffair(base.Affair)
	tagIDs, err = tagStruct.AddQuestionBoxTagByName(form.TagNames)
	if base.AffairRollbackIfError(err, c) {
		return
	}

	questionStruct.AffairInitWithAffair(base.Affair)
	err = questionStruct.Add()
	if base.AffairRollbackIfError(err, c) {
		return
	}
	err = tagStruct.AddQuestionBoxTagMap(questionStruct.QuestionBoxID, tagIDs)
	if base.AffairRollbackIfError(err, c) {
		return
	}
	if base.AffairFinished(c) {
		return
	}
	data := make(gin.H)
	data["question_box_id"] = questionStruct.QuestionBoxID
	response.OkWithData(data, c)
}


func LikeQuestionBox(c *gin.Context){
	var(
		err error
	)
	type LikeForm struct {
		QuestionBoxID int `json:"question_box_id"`
	}
	var form LikeForm
	err = c.ShouldBind(&form)
	if utils.FailOnError(e.INFO_ERROR, err, c){
		return
	}
	base := service.Base{}
	if base.AffairInit(c){
		return
	}
	base.AffairBegin()()

	questionBoxStruct := service.QuestionBoxStruct{QuestionBoxID:form.QuestionBoxID}


}