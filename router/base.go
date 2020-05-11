package router

import (
	v1 "finders-server/api/v1"
	"finders-server/global/response"

	"github.com/gin-gonic/gin"
)

//基本路由
func InitBaseRouter(Router *gin.RouterGroup) {

	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("register", v1.Register) //注册
		BaseRouter.GET("login", v1.Login)
	}
}

func record(c *gin.Context) {
	// response.Result(response.SUCCESS, []string{"elyar", "ablimit"}, "it is ok", c)
	response.OK(c)
}
