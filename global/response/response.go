/*
此模块用来封装response
*/

package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type WithToken struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Token string      `json:"token"`
}

const (
	ERROR                    = -1
	SUCCESS                  = 0 //定义当前的操作是成功的
	TOKEN_UPDATE_AND_SUCCESS = 1
)

func Result(code int, data interface{}, msg string, c *gin.Context) {

	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

//OK 当前操作成功用OK来回应
func OK(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

//OK 附加额外信息message
func OkWithMsg(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

//OK 附加data
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

// OK 附加token
func OKWithToken(token string, c *gin.Context) {
	c.JSON(http.StatusOK, WithToken{
		Code:  SUCCESS,
		Data:  "",
		Msg:   "操作成功",
		Token: token,
	})
}

//Fail 操作失败
func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

//Fail 附加信息
func FailWithMsg(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}
