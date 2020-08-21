package baseService

import (
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/pkg/e"
	"finders-server/utils/st"
	"github.com/gin-gonic/gin"
)

type Base struct {
	PageNum  int
	PageSize int
	Page     int

	Affair *model.AffairService
}

type AffairInterface interface {
	AffairInit(c *gin.Context) bool
	AffairBegin() func()
	AffairRollbackIfError(err error, c *gin.Context) bool
	AffairFinished(c *gin.Context) bool
	AffairInitWithAffair(a *model.AffairService)
	AffairRollback()
}

func (b *Base) AffairInit(context *gin.Context) bool {
	b.Affair = new(model.AffairService)
	err := b.Affair.NewAffairs()
	if err != nil {
		st.Debug(err)
		response.FailWithMsg(e.MYSQL_ERROR, context)
		return true
	}
	return false
}

func (b *Base) AffairInitWithAffair(a *model.AffairService) {
	b.Affair = new(model.AffairService)
	b.Affair.NewAffairsWithAffair(a)
}

func (b *Base) AffairBegin() func() {
	return b.Affair.DeferFunc()
}

func (b *Base) AffairRollbackIfError(err error, context *gin.Context) bool {
	if err != nil {
		b.Affair.RollBackIfError(err)
		st.Debug(err)
		response.FailWithMsg(e.MYSQL_ERROR, context)
		return true
	}
	return false
}

func (b *Base) AffairRollback() {
	b.Affair.RollBack()
}

func (b *Base) AffairFinished(context *gin.Context) bool {
	err := b.Affair.Commit()
	if err != nil {
		st.Debug(err)
		response.FailWithMsg(e.MYSQL_ERROR, context)
		return true
	}
	return false
}
