package middleware
<<<<<<< HEAD
=======

import (
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/pkg/e"
	"finders-server/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var user model.User
		var err error
		var jwtClaims *utils.JWTClaims
		code = e.Valid
		// 获取URL参数中的token
		token := c.GetHeader("token")
		if token == "" { // 若没有token则是不合法的参数
			code = e.TokenError
		} else { // 若有token则将token转换为 *Claims类型
			jwt := utils.NewJWT()
			// 解析token
			jwtClaims, err = jwt.ParseToken(token)
			if err != nil {
				code = e.TokenError
			} else if time.Now().Unix() > jwtClaims.ExpiredAt {
				// 若token过期了
				code = e.TokenOutOfDate
			}
		}

		// 若存在不合法 接下来的控制器或中间件就不需要执行了
		if code != e.Valid {
			response.FailWithMsg(e.GetMsg(code), c)
			c.Abort()
			return
		}
		if jwtClaims == nil{
			response.FailWithMsg(e.TOKEN_ERROR, c)
			return
		}
		user, err = model.GetUserByUserName(jwtClaims.UserName)
		if err != nil {
			response.FailWithMsg(e.MYSQL_ERROR, c)
			return
		}
		c.Request.Header.Set("username", user.UserName)
		c.Next()
	}
}
>>>>>>> test
