package requestForm

import (
	"finders-server/global/response"
	"finders-server/pkg/e"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type UpdateAdminForm struct {
	AdminName     string `json:"admin_name" validate:"omitempty,gte=1,lte=30"`
	AdminPassword string `json:"admin_password" validate:"omitempty,gte=1,lte=30"`
	Permission    int    `json:"permission" validate:"omitempty"`
}

func (u *UpdateAdminForm) Check(c *gin.Context) bool {
	validate := validator.New()
	err := validate.Struct(*u)
	if err != nil {
		response.FailWithMsg(e.INFO_ERROR, c)
		return true
	}
	return false
}
