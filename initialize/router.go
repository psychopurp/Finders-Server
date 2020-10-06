/*
初始化路由
*/

package initialize

import (
	"bufio"
	_ "finders-server/docs"
	"finders-server/global"
	"finders-server/global/response"
	"finders-server/middleware"
	"finders-server/router"
	"finders-server/service/upload"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {

	Router := gin.New()
	Router.Use(middleware.Cors())
	//加载静态资源
	Router.StaticFS("upload/images", http.Dir(upload.GetImageFullPath()))
	Router.StaticFS("upload/videos", http.Dir(upload.GetVideoFullPath()))
	Router.StaticFS("upload/fake", http.Dir("runtime/upload/fake/"))
	Router.LoadHTMLGlob("resource/*.html")
	Router.GET("", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	Router.GET("log", func(c *gin.Context) {
		file, err := os.Open("log/finders-server_info.log")
		if err != nil || file == nil {
			c.JSON(200, gin.H{"error": "error"})
		}
		defer file.Close()
		br := bufio.NewReader(file)
		var Listdata []string
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			Listdata = append(Listdata, string(a))
		}
		c.HTML(200, "log.html", gin.H{"Listdata": Listdata})
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
	router.InitMomentRouter(APIGroup)
	router.InitRecommendRouter(APIGroup)
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
