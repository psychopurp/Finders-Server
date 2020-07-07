package router

import (
	v1 "finders-server/api/v1"
	"finders-server/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdminRouter(Router *gin.RouterGroup) {

	AdminRouter := Router.Group("admin")
	{
		AdminRouter.POST("login", v1.AdminLogin)                                         // 登陆
		AdminRouter.POST("update_profile", middleware.AdminJWT(), v1.UpdateAdminProfile) // 登陆
	}
}
