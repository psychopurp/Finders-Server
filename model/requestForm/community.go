package requestForm

import (
	"finders-server/global/response"
	"finders-server/pkg/e"
	"finders-server/st"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type CreateCommunityForm struct {
	CommunityName        string `json:"community_name" validate:"required,min=1,max=100"`
	CommunityDescription string `json:"community_description" validate:"required,min=1,max=65535"`
	Background           string `json:"background" validate:"required,min=1,max=200"`
	CommunityAvatar      string `json:"community_avatar" validate:"required,min=1,max=200"`
}

func (u *CreateCommunityForm) Check(c *gin.Context) bool {
	validate := validator.New()
	err := validate.Struct(*u)
	if err != nil {
		st.DebugWithFuncName(err)
		response.FailWithMsg(e.INFO_ERROR, c)
		return true
	}
	return false
}

type UpdateCommunityForm struct {
	CommunityID          int    `json:"community_id" validate:"required,gte=0"`
	CommunityName        string `json:"community_name" validate:"omitempty,min=1,max=100"`
	CommunityDescription string `json:"community_description" validate:"omitempty,min=1,max=65535"`
	Background           string `json:"background" validate:"omitempty,min=1,max=200"`
	CommunityAvatar      string `json:"community_avatar" validate:"required,min=1,max=200"`
}

func (u *UpdateCommunityForm) Check(c *gin.Context) bool {
	validate := validator.New()
	err := validate.Struct(*u)
	if err != nil {
		st.DebugWithFuncName(err)
		response.FailWithMsg(e.INFO_ERROR, c)
		return true
	}
	return false
}

type GetCommunityIDForm struct {
	CommunityID int `json:"community_id" validate:"required,gte=0"`
}

func (u *GetCommunityIDForm) Check(c *gin.Context) bool {
	validate := validator.New()
	err := validate.Struct(*u)
	if err != nil {
		st.DebugWithFuncName(err)
		response.FailWithMsg(e.INFO_ERROR, c)
		return true
	}
	return false
}
