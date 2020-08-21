package router

import (
	v1 "finders-server/api/v1"
	"finders-server/middleware"
	"github.com/gin-gonic/gin"
)

func InitActivityRouter(Router *gin.RouterGroup) {

	ActivityRouter := Router.Group("activity")
	{
		ActivityRouter.GET("get_activities", v1.GetActivities)
		ActivityRouter.GET("get_user_activities", v1.GetUserActivities)
		ActivityRouter.POST("add_activity", middleware.JWT(), v1.AddActivity)
		ActivityRouter.GET("get_activity_info", v1.GetActivityInfo)
		ActivityRouter.POST("collect", middleware.JWT(), v1.CollectActivity)
		ActivityRouter.POST("uncollect", middleware.JWT(), v1.UnCollectActivity)
		ActivityRouter.GET("get_collect", middleware.JWT(), v1.GetCollectActivities)
		ActivityRouter.GET("get_activity_like", middleware.JWT(), v1.GetActivityLike)
		ActivityRouter.POST("like_activity", middleware.JWT(), v1.LikeActivity)
		ActivityRouter.POST("dislike_activity", middleware.JWT(), v1.DisLikeActivity)
		ActivityRouter.POST("comment", middleware.JWT(), v1.AddComment)
		ActivityRouter.POST("reply", middleware.JWT(), v1.AddReply)
		ActivityRouter.GET("get_activity_comment", v1.GetActivityComments)
		ActivityRouter.GET("get_comment_reply", v1.GetCommentReplies)

	}
}
