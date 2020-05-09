package service

import (
	"errors"
	"finders-server/global"
	"finders-server/model"
	"finders-server/utils"

	uuid "github.com/satori/go.uuid"
)

/*
用户相关的一些service
*/

// @title    Register
// @description   register, 用户注册
// @param     u               *model.User
// @return    err             error
// @return    user       *model.User

func Register(u model.User) (model.User, error) {
	var user model.User
	db := global.DB
	//判断用户是否注册
	isRegister := !db.Where("username = ?", u.UserName).First(&user).RecordNotFound()
	if isRegister {
		return user, errors.New("用户名已注册")
	} else {
		//给用户用户ID进行注册
		u.UserID = uuid.NewV4()
		//将密码加密
		u.Password = utils.MD5(u.Password)
		err := db.Create(&u).Error
		return u, err

	}

}
