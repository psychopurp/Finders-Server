package requestForm

import (
	"finders-server/global/response"
	"finders-server/pkg/e"
	"finders-server/st"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type CreateCommentForm struct {
	ItemID  string `json:"item_id" validate:"required,min=1,max=50"`
	Content string `json:"content" validate:"required,min=1,max=65535"`
}

func (u *CreateCommentForm) Check(c *gin.Context) bool {
	validate := validator.New()
	err := validate.Struct(*u)
	if err != nil {
		st.DebugWithFuncName(err)
		response.FailWithMsg(e.INFO_ERROR, c)
		return true
	}
	return false
}
