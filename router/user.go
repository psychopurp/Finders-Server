package router

import (
	v1 "finders-server/api/v1"
	"github.com/gin-gonic/gin"
)

//基本路由
func InitUserRouter(Router *gin.RouterGroup) {

	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", v1.Login)                    // 登陆
		UserRouter.POST("update_profile", v1.UpdateProfile)   // 用户信息更新
		UserRouter.POST("follow", v1.Follow)                  // 关注用户
		UserRouter.POST("unfollow", v1.UnFollow)              // 取消关注用户
		UserRouter.POST("add_denylist", v1.AddDenyList)       // 添加黑名单
		UserRouter.POST("remove_denylist", v1.RemoveDenyList) // 移除黑名单
		UserRouter.GET("get_denylist", v1.GetDenyList)        // 查看黑名单
		UserRouter.GET("get_fans", v1.GetFans)                // 查看用户的粉丝列表
		UserRouter.GET("get_follow", v1.GetFollow)            // 查看用户的关注
	}
}
