/*
初始化路由
*/

package initialize

import (
	_ "finders-server/docs"
	"finders-server/global"
	"finders-server/global/response"
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
	Router.LoadHTMLGlob("resource/*.html")
	Router.GET("", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
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

// @获取指定ID记录
// @Description get record by ID
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "userId"
// @Success 200 {string} string "{"code": 0 ,"data":{} ,"msg":""}"
// @Router /hello [get]
func record(c *gin.Context) {
	// response.Result(response.SUCCESS, []string{"elyar", "ablimit"}, "it is ok", c)
	response.OK(c)
}
