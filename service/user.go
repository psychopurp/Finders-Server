package service

import (
	"errors"
	"finders-server/model"
	"finders-server/model/requestForm"
	"finders-server/model/responseForm"
	"finders-server/pkg/e"
	"finders-server/utils"
	uuid "github.com/satori/go.uuid"
	"math"
	"time"
)

/*
用户相关的一些service
*/

type UserStruct struct {
	UserID   uuid.UUID
	Phone    string
	Password string
	Nickname string
	Status   int
	Avatar   string
	UserInfo UserInfoStruct
	UserName string

	PageNum  int
	PageSize int
	Page     int
}

type UserInfoStruct struct {
	TrueName      string
	Address       string
	Sex           string
	Sexual        string
	Feeling       string
	Birthday      string
	Introduction  string
	Signature     string
	BloodType     string
	Eamil         string
	QQ            string
	Wechat        string
	Profession    string
	School        string
	Constellation string
	Credit        int
	UserTag       string
	Age           int
}

func (u *UserStruct) ExistUserByUserNameAndPassword() (user model.User, isExist bool) {
	user, isExist = model.ExistUserByUserNameAndPassword(u.UserName, u.Password)
	if isExist {
		u.UserID = user.UserID
		u.UserName = user.UserName
	}
	return
}

func (u *UserStruct) GetUserByUserName() (user model.User, err error) {
	return model.GetUserByUserName(u.UserName)
}

func (u *UserStruct) ExistUserByPhone() (user model.User, isExist bool) {
	user, isExist = model.ExistUserByPhone(u.Phone)
	if isExist {
		u.UserName = user.UserName
		u.UserID = user.UserID
	}
	return
}

func (u *UserStruct) Register() (user model.User, err error) {
	data := map[string]interface{}{
		"phone":    u.Phone,
		"user_id":  uuid.NewV4(),
		"username": uuid.NewV4().String(),
		"nickname": uuid.NewV4().String()[:20],
		"password": uuid.NewV4().String()[:20],
		"status":   model.Normal,
	}

	user, err = model.AddUserByMap(data)
	if err != nil {
		return
	}
	u.UserID = user.UserID
	u.UserName = user.UserName
	return
}

// 通过用户名获取 对应的token
func (u *UserStruct) GetAuth() (token string, err error) {
	jwt := utils.NewJWT()
	createAt := time.Now()
	expiredAt := createAt.Add(time.Hour * 5)
	jwtClaims := utils.JWTClaims{
		MapClaims: nil,
		UserID:    u.UserID.String(),
		CreatedAt: createAt.Unix(),
		ExpiredAt: expiredAt.Unix(),
	}
	token, err = jwt.GenerateToken(jwtClaims)
	if err != nil {
		return
	}
	model.AddLoginRecord(u.UserID.String(), token, createAt, expiredAt)
	return
}

func (u *UserStruct) BindUpdateForm(form requestForm.UserUpdateForm) (err error) {
	u.Password = form.Password
	u.UserName = form.UserName
	u.Avatar = form.Avatar
	u.Nickname = form.Nickname
	u.UserInfo = UserInfoStruct{
		TrueName:      form.TrueName,
		Address:       form.Address,
		Sex:           form.Sex,
		Sexual:        form.Sexual,
		Feeling:       form.School,
		Birthday:      form.School,
		Introduction:  form.Introduction,
		BloodType:     form.BloodType,
		Eamil:         form.Eamil,
		QQ:            form.QQ,
		Wechat:        form.Wechat,
		Profession:    form.Profession,
		School:        form.School,
		Constellation: form.Constellation,
		Credit:        form.Credit,
		Age:           form.Age,
		Signature:     form.Signature,
	}
	return nil
}

func (u *UserStruct) UpdateUserInfo() (err error) {
	userInfo := model.UserInfo{
		TrueName:      u.UserInfo.TrueName,
		Address:       u.UserInfo.Address,
		Sex:           u.UserInfo.Sex,
		Sexual:        u.UserInfo.Sexual,
		Feeling:       u.UserInfo.Feeling,
		Birthday:      u.UserInfo.Birthday,
		Introduction:  u.UserInfo.Introduction,
		BloodType:     u.UserInfo.BloodType,
		Eamil:         u.UserInfo.Eamil,
		QQ:            u.UserInfo.QQ,
		Wechat:        u.UserInfo.Wechat,
		Profession:    u.UserInfo.Profession,
		School:        u.UserInfo.School,
		Constellation: u.UserInfo.Constellation,
		Credit:        u.UserInfo.Credit,
		Age:           u.UserInfo.Age,
		Signature:     u.UserInfo.Signature,
	}
	return model.UpdateUserInfoByUserInfo(u.UserID.String(), userInfo)
}

func (u *UserStruct) GetUserByUserID() (user model.User, err error) {
	return model.GetUserByUserID(u.UserID.String())
}
func (u *UserStruct) UpdateUser() (err error) {
	user := model.User{
		Password: u.Password,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		UserName: u.UserName,
	}
	return model.UpdateUserByUser(u.UserID.String(), user)
}

func (u *UserStruct) Edit() (err error) {
	err = u.UpdateUser()
	if err != nil {
		return
	}
	err = u.UpdateUserInfo()
	return
}

func (u *UserStruct) CheckExistRelation(toID string, relationType int) (ok bool) {
	ok, _ = model.ExistRelation(u.UserID.String(), toID, relationType)
	return
}

func (u *UserStruct) AddRelation(toID string, relationType int) (relation model.Relation, err error) {
	var (
		isExist bool
	)
	// 若想添加 黑名单关系
	if relationType == model.DENY {
		// 看看两者有没有关注关系
		isExist, err = model.ExistRelation(u.UserID.String(), toID, model.FOLLOW)
		if err != nil {
			return relation, errors.New(e.MYSQL_ERROR)
		}
		// 若有 可以直接修改关注为拉黑
		if isExist {
			data := make(map[string]interface{})
			data["from_uid"] = u.UserID.String()
			data["to_uid"] = toID
			data["relation_type"] = model.FOLLOW
			relation, err = model.UpdateRelationType(data, model.DENY)
			return
		}
	} else {
		// 若是要添加关注关系
		// 先看是否已存在关注关系
		isExist, err = model.ExistRelation(u.UserID.String(), toID, relationType)
		if err != nil {
			return relation, errors.New(e.MYSQL_ERROR)
		}
		// 若存在则
		if isExist {
			return
		}

	}
	relation, err = model.AddRelation(u.UserID.String(), toID, relationType)
	if err != nil {
		return relation, errors.New(e.MYSQL_ERROR)
	}
	return
}

func (u *UserStruct) DeleteRelation(toID string, relationType int) (relation model.Relation, err error) {
	var exist bool
	exist, err = model.ExistRelation(u.UserID.String(), toID, relationType)
	if err != nil || !exist {
		return
	}
	relation, err = model.DeleteRelation(u.UserID.String(), toID, relationType)
	return
}

const (
	FROM = 1
	TO   = 2
)

func (u *UserStruct) GetRelationsWithPageByData(data map[string]interface{}) (relations []*model.Relation, err error) {
	return model.GetRelationsWithPageByData(data, u.PageNum, u.PageSize)
}

func (u *UserStruct) GetSimpleUserInfoListWitPageByUserID(relationType, loc int) (form responseForm.SimpleUserInfoWithPage, err error) {
	var (
		relations []*model.Relation
		users     []*model.User
		totalCNT  int
		//userInfo  model.UserInfo
	)
	form.Page = u.Page
	data := make(map[string]interface{})
	data["relation_type"] = relationType
	// 根据粉丝列表还是关注列表选择 userID是属于from还是to
	if loc == FROM {
		data["from_uid"] = u.UserID.String()
	} else {
		data["to_uid"] = u.UserID.String()
	}
	totalCNT, err = model.GetRelationTotalByData(data)
	form.TotalCNT = totalCNT
	if form.TotalCNT == 0 {
		return
	}
	relations, err = u.GetRelationsWithPageByData(data)
	if err != nil {
		return
	}
	form.CNT = len(relations)
	form.TotalPage = int(math.Ceil(float64(totalCNT) / float64(u.PageSize)))
	var userIDs []string
	for _, relation := range relations {
		if loc == FROM {
			userIDs = append(userIDs, relation.ToUID.String())
		} else {
			userIDs = append(userIDs, relation.FromUID.String())
		}
	}
	users, err = model.GetUsersByUserIDs(userIDs)
	for _, user := range users {
		if err != nil {
			return form, errors.New(e.MYSQL_ERROR)
		}
		form.SimpleUserInfos = append(form.SimpleUserInfos, responseForm.SimpleUserInfo{
			UserId:       user.UserID.String(),
			Avatar:       user.Avatar,
			NickName:     user.Nickname,
			Introduction: user.UserInfo.Introduction,
			Signature:    user.UserInfo.Signature,
		})
	}
	return

}

//func (u *UserStruct) GetSimpleUserInfoListByUserID(relationType int, loc int) (simpleUserInfo []responseForm.SimpleUserInfo, err error) {
//	var (
//		isExist   bool
//		relations []model.Relation
//		users     []model.User
//		//userInfo  model.UserInfo
//	)
//	data := make(map[string]interface{})
//	data["relation_type"] = relationType
//	// 根据粉丝列表还是关注列表选择 userID是属于from还是to
//	if loc == FROM {
//		data["from_uid"] = u.UserID.String()
//	} else {
//		data["to_uid"] = u.UserID.String()
//	}
//	isExist, err = model.ExistRelationByData(data)
//	if err != nil {
//		return
//	}
//	if !isExist {
//		return
//	}
//	relations, err = model.GetRelationsByData(data)
//	if err != nil {
//		return
//	}
//	var userIDs []string
//	for _, relation := range relations {
//		if loc == FROM {
//			userIDs = append(userIDs, relation.ToUID.String())
//		} else {
//			userIDs = append(userIDs, relation.FromUID.String())
//		}
//	}
//	users, err = model.GetUsersByUserIDs(userIDs)
//	for _, user := range users {
//		//userInfo, err = model.GetUserInfoByUserID(user.UserID.String())
//		if err != nil {
//			return simpleUserInfo, errors.New(e.MYSQL_ERROR)
//		}
//		simpleUserInfo = append(simpleUserInfo, responseForm.SimpleUserInfo{
//			UserId:   user.UserID.String(),
//			Avatar:   user.Avatar,
//			NickName: user.Nickname,
//			//Introduction: userInfo.Introduction,
//			Introduction: user.UserInfo.Introduction,
//			Signature:    user.UserInfo.Signature,
//		})
//	}
//	return
//}
