package router

import (
	v1 "finders-server/api/v1"
	"finders-server/middleware"
	"github.com/gin-gonic/gin"
)

// user路由
func InitUserRouter(Router *gin.RouterGroup) {

	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", v1.Login)                                      // 登陆
		UserRouter.POST("update_profile", middleware.JWT(), v1.UpdateProfile)   // 用户信息更新
		UserRouter.GET("user_info", middleware.JWT(), v1.GetUserInfo)   	// 获取用户信息
		UserRouter.POST("follow", middleware.JWT(), v1.Follow)                  // 关注用户
		UserRouter.POST("unfollow", middleware.JWT(), v1.UnFollow)              // 取消关注用户
		UserRouter.POST("add_denylist", middleware.JWT(), v1.AddDenyList)       // 添加黑名单
		UserRouter.POST("remove_denylist", middleware.JWT(), v1.RemoveDenyList) // 移除黑名单
		UserRouter.GET("get_denylist", middleware.JWT(), v1.GetDenyList)        // 查看黑名单
		UserRouter.GET("get_followlist", middleware.JWT(), v1.GetFollowList)    // 查看黑名单
		UserRouter.GET("get_fans", v1.GetFans)                                  // 查看用户的粉丝列表
		UserRouter.GET("get_follow", v1.GetFollow)                              // 查看用户的关注
		UserRouter.GET("check_follow", middleware.JWT(), v1.CheckFollow)        // 查看用户是否关注某人
	}
}
