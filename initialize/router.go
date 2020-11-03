/*
初始化路由
*/

package initialize

import (
	_ "finders-server/docs"
	"finders-server/global"
	"finders-server/middleware"
	"finders-server/router"
	"finders-server/service/upload"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {

	Router := gin.New()
	//加载静态资源
	Router.StaticFS("upload/images", http.Dir(upload.GetImageFullPath()))
	Router.StaticFS("upload/videos", http.Dir(upload.GetVideoFullPath()))
	// Router.LoadHTMLGlob("resource/*.html")
	Router.GET("", func(c *gin.Context) {
		c.Header("Content-Type", "text/html;charset=utf-8")
		c.String(200, "<h1>Hello World</h1>")
	})

	Router.Use(middleware.Logger())
	global.LOG.Debug("Register Logger middleware")

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Debug("Register swagger handler")

	// 方便统一添加路由组前缀 多服务器上线使用
	APIGroup := Router.Group("/api/v1")
	router.InitBaseRouter(APIGroup) //注册基本路由 不用鉴权
	router.InitUserRouter(APIGroup)
	router.InitAdminRouter(APIGroup)
	router.InitMediaRouter(APIGroup)
	router.InitCommunityRouter(APIGroup)
	router.InitActivityRouter(APIGroup)
	return Router

}
