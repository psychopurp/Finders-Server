package requestForm

import (
	"finders-server/global/response"
	"finders-server/pkg/e"
	"finders-server/st"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type CreateMomentRequestForm struct {
	MomentInfo string   `json:"moment_info" validate:"required,min=1,max=500"`
	MediaIDs   []string `json:"media_ids"  validate:"required,min=1,max=500"`
	Location   string   `json:"location" validate:"required,min=1,max=500"`
}

func (u *CreateMomentRequestForm) Check(c *gin.Context) bool {
	validate := validator.New()
	err := validate.Struct(*u)
	if err != nil {
		st.DebugWithFuncName(err)
		response.FailWithMsg(e.INFO_ERROR, c)
		return true
	}
	return false
}

type LikeMomentRequestForm struct {
	MomentID string `json:"moment_info" validate:"required,min=1,max=50"`
}

func (u *LikeMomentRequestForm) Check(c *gin.Context) bool {
	validate := validator.New()
	err := validate.Struct(*u)
	if err != nil {
		st.DebugWithFuncName(err)
		response.FailWithMsg(e.INFO_ERROR, c)
		return true
	}
	return false
}
