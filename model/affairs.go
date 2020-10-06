package model

import (
	"finders-server/global"
	"github.com/jinzhu/gorm"
)

type AffairService struct {
	tx *gorm.DB
}

func (a *AffairService) NewAffairs() (err error) {
	db := global.DB
	a.tx = db.Begin()
	if err = a.tx.Error; err != nil {
		return
	}
	return nil
}
func (a *AffairService) NewAffairsWithAffair(n *AffairService) {
	a.tx = n.tx
}
func (a *AffairService) DeferFunc() func() {
	if a.tx == nil {
		return func() {
			panic("affair not init")
		}
	}
	return func() {
		if r := recover(); r != nil {
			a.tx.Rollback()
		}
	}
}

func (a *AffairService) RollBackIfError(err error) {
	if err != nil {
		a.tx.Rollback()
	}
}
func (a *AffairService) RollBack() {
	a.tx.Rollback()
}
func (a *AffairService) Commit() (err error) {
	return a.tx.Commit().Error
}

func (a *AffairService) GetTX() *gorm.DB {
	return a.tx
}

// 在每个service中都加上 affairService
// 所有model.Addxxx等操作都变成affairService的函数
