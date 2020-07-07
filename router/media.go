package router

import (
	v1 "finders-server/api/v1"
	"finders-server/middleware"
	"github.com/gin-gonic/gin"
)

func InitMediaRouter(Router *gin.RouterGroup) {

	AdminRouter := Router.Group("media")
	{
		AdminRouter.POST("upload_image", middleware.JWT(), v1.UploadImage)
		AdminRouter.POST("upload_video", middleware.JWT(), v1.UploadVideo)
	}
}
