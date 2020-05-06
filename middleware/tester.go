package middleware

import (
	"finders-server/global"

	"github.com/gin-gonic/gin"
)

func Tester() gin.HandlerFunc {

	return func(c *gin.Context) {
		//开始时间

		global.LOG.Debug("处理Tester Next前")
		//处理请求
		c.Next()

		global.LOG.Debug("处理Tester next后")

	}

}
