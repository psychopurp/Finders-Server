package utils

import (
	"errors"
	"finders-server/global"
	"finders-server/global/response"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetErrorAndLog(errStr string, err error, funcName string) error {
	if errStr == "" {
		global.LOG.Debug(err.Error(), " func: ", funcName)
		return err
	}
	if err != nil {
		global.LOG.Debug(err.Error(), " func: ", funcName)
	} else {
		global.LOG.Debug(errStr, " func: ", funcName)
	}
	return errors.New(errStr)
}

// GetPage get page parameters
func GetPage(c *gin.Context) (pageNum int, page int) {
	pageNum = 0
	page = com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	if page > 0 {
		pageNum = (page - 1) * global.CONFIG.AppSetting.PageSize
	}
	return
}

func FailOnError(errStr string, err error, c *gin.Context) bool {
	if err != nil {
		global.LOG.Debug(err.Error())
		if errStr != "" {
			response.FailWithMsg(errStr, c)
		} else {
			response.FailWithMsg(err.Error(), c)
		}
		return true
	}
	return false
}
