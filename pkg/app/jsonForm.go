package app

import (
	"errors"
	"finders-server/pkg/e"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// usage: BindAndValid(c, form)
func BindAndValid(c *gin.Context, json interface{}) (err error) {
	err = c.BindJSON(&json)
	if err != nil {
		return errors.New(e.INFO_ERROR)
	}
	validate := validator.New()
	err = validate.Struct(json)
	if err != nil {
		return errors.New(e.INFO_ERROR)
	}
	return
}
