package router

import (
	v1 "finders-server/api/v1"
	"finders-server/middleware"
	"github.com/gin-gonic/gin"
)

func InitCommunityRouter(Router *gin.RouterGroup) {
	CommunityRouter := Router.Group("community")
	{
		CommunityRouter.POST("create_community", middleware.JWT(), v1.CreateCommunity)
		CommunityRouter.POST("update_profile", middleware.JWT(), v1.UpdateCommunityProfile)
		CommunityRouter.POST("collect", middleware.JWT(), v1.CollectCommunity)
		CommunityRouter.POST("uncollect", middleware.JWT(), v1.UnCollectCommunity)
		CommunityRouter.GET("get_collect", middleware.JWT(), v1.GetCollectCommunity)
	}
}
