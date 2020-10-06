package requestForm

import (
	"finders-server/global/response"
	"finders-server/pkg/e"
	"finders-server/st"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type CreateActivityForm struct {
	CommunityID   int      `json:"community_id" validate:"required,gte=0"`
	ActivityInfo  string   `json:"activity_info" validate:"required,min=1,max=65535"`
	ActivityTitle string   `json:"activity_title" validate:"required,min=1,max=50"`
	MediaIDs      []string `json:"media_ids" validate:"required,min=1,max=50"`
}

type GetActivityIDForm struct {
	ActivityID string `json:"activity_id" validate:"required,min=1,max=50"`
}

func (u *CreateActivityForm) Check(c *gin.Context) bool {
	validate := validator.New()
	err := validate.Struct(*u)
	if err != nil {
		st.DebugWithFuncName(err)
		response.FailWithMsg(e.INFO_ERROR, c)
		return true
	}
	return false
}
