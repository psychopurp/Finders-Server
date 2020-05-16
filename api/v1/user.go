package v1

import (
	"finders-server/global/response"

	"github.com/gin-gonic/gin"
)

/*
用户相关接口
*/

func Register(c *gin.Context) {

	response.OkWithMsg("用户注册  test", c)
}

func Login(c *gin.Context) {
	response.OkWithMsg("用户登陆", c)
}
