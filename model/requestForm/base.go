package requestForm

import "github.com/gin-gonic/gin"

type CheckInterface interface {
	Check(c *gin.Context) bool
}
