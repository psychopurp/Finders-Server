package router

import (
	v1 "finders-server/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRecommendRouter(Router *gin.RouterGroup) {

	ReommendRouter := Router.Group("recommend")
	{
		ReommendRouter.GET("main", v1.MainRecommend)
	}
}
