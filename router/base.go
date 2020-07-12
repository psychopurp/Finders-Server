package router

import (
<<<<<<< HEAD
	v1 "finders-server/api/v1"

=======
>>>>>>> test
	"github.com/gin-gonic/gin"
)

//基本路由
func InitBaseRouter(Router *gin.RouterGroup) {

<<<<<<< HEAD
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("register", v1.Register) //注册
		BaseRouter.GET("login", v1.Login)
	}
=======
	//BaseRouter := Router.Group("base")

		//BaseRouter.GET("register", v1.Register) //注册
		//BaseRouter.GET("login", v1.Login)

>>>>>>> test
}
