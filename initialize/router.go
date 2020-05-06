/*
初始化路由
*/

package initialize

import (
	_ "finders-server/docs"
	"finders-server/global"
	"finders-server/global/response"
	"finders-server/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {

	Router := gin.New()

	global.LOG.Debug("Create gin Engine")
	Router.Use(middleware.Logger())
	global.LOG.Debug("Register Logger middleware")

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Debug("Register swagger handler")

	Router.GET("/hello", record)
	Router.GET("/", func(c *gin.Context) {
		response.OK(c)

	})
	// Router.Run(":8550")
	return Router

}

// @获取指定ID记录
// @Description get record by ID
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "userId"
// @Success 200 {string} string	"ok"
// @Router /record/{some_id} [get]
func record(c *gin.Context) {
	// response.Result(response.SUCCESS, []string{"elyar", "ablimit"}, "it is ok", c)
	response.OK(c)
}
