package middleware

import (
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/pkg/cache"
	"finders-server/pkg/e"
	"finders-server/st"
	"finders-server/utils"
	"fmt"
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
		if jwtClaims == nil {
			response.FailWithMsg(e.TOKEN_ERROR, c)
			c.Abort()
			return
		}
		cacheSrv := cache.NewUserCacheService()
		user, err = cacheSrv.GetUserByUserId(jwtClaims.UserID)
		if err != nil {
			user, err = model.GetUserByUserID(jwtClaims.UserID)
			err = cacheSrv.SetUserByUserId(user)
			if err != nil {
				st.Debug("jwt set cache error", err)
			}
		}
		if err != nil {
			response.FailWithMsg(e.MYSQL_ERROR, c)
			c.Abort()
			return
		}
		c.Request.Header.Set("username", user.UserName)
		c.Request.Header.Set("user_id", jwtClaims.UserID)
		c.Next()
	}
}

func JWTOmitEmpty() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var user model.User
		var err error
		var jwtClaims *utils.JWTClaims
		code = e.Valid
		// 获取URL参数中的token
		token := c.GetHeader("token")
		if token != "" { // 若有token则将token转换为 *Claims类型
			jwt := utils.NewJWT()
			// 解析token
			jwtClaims, err = jwt.ParseToken(token)
			if err != nil {
				code = e.TokenError
			} else if time.Now().Unix() > jwtClaims.ExpiredAt {
				// 若token过期了
				code = e.TokenOutOfDate
			}
			// 若存在不合法 接下来的控制器或中间件就不需要执行了
			if code != e.Valid {
				response.FailWithMsg(e.GetMsg(code), c)
				c.Abort()
				return
			}
			if jwtClaims == nil {
				response.FailWithMsg(e.TOKEN_ERROR, c)
				c.Abort()
				return
			}
			user, err = model.GetUserByUserID(jwtClaims.UserID)
			if err != nil {
				response.FailWithMsg(e.MYSQL_ERROR, c)
				c.Abort()
				return
			}
			c.Request.Header.Set("username", user.UserName)
			c.Request.Header.Set("user_id", jwtClaims.UserID)
		}

		c.Next()
	}
}
func AdminJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var admin model.Admin
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
		if jwtClaims == nil {
			response.FailWithMsg(e.TOKEN_ERROR, c)
			c.Abort()
			return
		}
		fmt.Println(jwtClaims.UserID)
		admin, err = model.GetAdminByAdminID(jwtClaims.UserID)
		if err != nil {
			response.FailWithMsg(e.MYSQL_ERROR, c)
			c.Abort()
			return
		}
		c.Request.Header.Set("adminName", admin.AdminName)
		c.Request.Header.Set("admin_id", admin.AdminID)
		c.Next()
	}
}
