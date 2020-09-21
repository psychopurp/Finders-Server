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
		MomentRouter.GET("get_user_moments", middleware.JWT(), v1.GetUserMoments)
		MomentRouter.GET("get_moment", middleware.JWT(), v1.GetMoment)
		MomentRouter.GET("like_moment", middleware.JWT(), v1.LikeMoment)
	}
}
