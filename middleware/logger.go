package middleware

import (
	"finders-server/global"
	"time"

	"github.com/gin-gonic/gin"
)

//Logger 是日志中间件
func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()

		// global.LOG.Debug("处理Next前")
		//处理请求
		c.Next()

		// global.LOG.Debug("处理next后")

		endTime := time.Now()

		//执行时间
		latencyTime := endTime.Sub(startTime)

		//请求方式
		reqMethod := c.Request.Method

		//请求路由
		reqURI := c.Request.RequestURI

		//状态码
		statusCode := c.Writer.Status()

		//请求IP
		clientIP := c.ClientIP()

		// 日志格式
		global.LOG.Infof("| status = %d | exec = %v | clientIP = %s | method = %s | URI = %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqURI,
		)

	}

}
