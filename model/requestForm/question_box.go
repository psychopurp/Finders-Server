package requestForm

import (
	"finders-server/global/response"
	"finders-server/pkg/e"
	"finders-server/st"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type CreateQuestionBoxForm struct {
	QuestionBoxInfo string   `json:"question_box_info" validate:"required"`
	TagNames        []string `json:"tag_names"`
}

func (u *CreateQuestionBoxForm) Check(c *gin.Context) bool {
	validate := validator.New()
	err := validate.Struct(*u)
	if err != nil {
		st.DebugWithFuncName(err)
		response.FailWithMsg(e.INFO_ERROR, c)
		return true
	}
	return false
}
