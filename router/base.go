package router

import (
	v1 "finders-server/api/v1"

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
