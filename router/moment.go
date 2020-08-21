package router

import (
	v1 "finders-server/api/v1"
	"finders-server/middleware"
	"github.com/gin-gonic/gin"
)

func InitMomentRouter(Router *gin.RouterGroup) {

	MomentRouter := Router.Group("moment")
	{
		MomentRouter.POST("create_moment", middleware.JWT(), v1.CreateMoment)
		MomentRouter.GET("get_moments", middleware.JWT(), v1.GetMoments)
	}
}
