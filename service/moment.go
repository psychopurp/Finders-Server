package service

import (
	"finders-server/model"
	"finders-server/model/responseForm"
	"finders-server/pkg/e"
	"finders-server/utils"
	"math"
	"strings"
)

type MomentStruct struct {
	MomentID     string
	MomentStatus int
	MomentInfo   string
	ReadNum      int // 曝光数
	//ActivityTag    string `gorm:"column:activity_tag;type:TEXT;size:65535;" json:"activity_tag"`          //[ 7] activity_tag                                   TEXT[65535]          null: true   primary: false  auto: false
	MediaIDs string
	Location string
	//MediaType   int    `gorm:"column:media_type;int;" json:"media_id"`       //[ 8] picture_id                                     VARCHAR[30]          null: true   primary: false  auto: false
	UserID string

	Base
}

func (m *MomentStruct) Add() (moment model.Moment, err error) {
	data := map[string]interface{}{
		"moment_status": model.MomentNormal,
		"moment_info":   m.MomentInfo,
		"media_ids":     m.MediaIDs,
		"location":      m.Location,
		"user_id":       m.UserID,
	}
	moment, err = m.Affair.AddMomentByMap(data)
	return
}

func (m *MomentStruct) GetMomentsTotalByUserID() (cnt int, err error) {
	return model.GetMomentsTotal(m.UserID)
}

func (m *MomentStruct) GetMomentsByUserID() (moments []*model.Moment, err error) {
	return model.GetMomentsWithPaginator(m.PageNum, m.PageSize, m.UserID)
}

func (m *MomentStruct) GetMomentByMomentID() (moment model.Moment, err error) {
	return model.GetMomentByMomentID(m.MomentID)
}

func (m *MomentStruct) AddReadNum() (err error) {
	return m.Affair.AddMomentReadNum(m.MomentID)
}

func (m *MomentStruct) GetUserMomentsResponseForm() (form responseForm.GetUserMomentsResponseForm, err error) {
	var (
		moments  []*model.Moment
		totalCnt int
		user     model.User
	)
	form.Page = m.Page
	totalCnt, _ = m.GetMomentsTotalByUserID()
	form.TotalCNT = totalCnt
	if totalCnt == 0 {
		return
	}
	form.TotalPage = int(math.Ceil(float64(totalCnt) / float64(m.PageSize)))
	moments, err = m.GetMomentsByUserID()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "GetActivitiesPageResponse1")
		return
	}
	form.CNT = len(moments)
	user, err = model.GetUserByUserID(m.UserID)
	if err != nil {
		return
	}
	form.UserID = m.UserID
	form.Avatar = user.Avatar
	form.NickName = user.Nickname
	for _, moment := range moments {
		var mediasForms []responseForm.MediasForm
		if len(moment.MediaIDs) != 0 {
			mediaIDs := strings.Split(moment.MediaIDs, ";")
			for _, mediaID := range mediaIDs {
				var media model.Media
				media, err = model.GetMediaByMediaID(mediaID)
				if err != nil {
					return
				}
				mediaForm := responseForm.MediasForm{
					MediaURL:  media.MediaURL,
					MediaType: model.GetMediaTypeByInt(media.MediaType),
				}
				mediasForms = append(mediasForms, mediaForm)
			}
		}
		m.MomentID = moment.MomentID
		err = m.AddReadNum()
		if err != nil {
			return
		}
		//simpleForm := responseForm.SimpleMomentForm{
		//	MomentID:   moment.MomentID,
		//MomentInfo: moment.MomentInfo,
		//ReadNum:    moment.ReadNum + 1,
		//Medias:     mediasForms,
		//CreatedAt:  moment.CreatedAt.String(),
		//}
		form.MomentIDs = append(form.MomentIDs, moment.MomentID)
	}
	return
}

func (m *MomentStruct) GetMomentResponseForm() (form responseForm.GetMomentResponseForm, err error) {
	var (
		moment      model.Moment
		user        model.User
		mediasForms []responseForm.MediasForm
	)
	moment, err = m.GetMomentByMomentID()
	if err != nil {
		return
	}
	user, err = model.GetUserByUserID(moment.UserID)
	if err != nil {
		return
	}
	if len(moment.MediaIDs) != 0 {
		mediaIDs := strings.Split(moment.MediaIDs, ";")
		for _, mediaID := range mediaIDs {
			var media model.Media
			media, err = model.GetMediaByMediaID(mediaID)
			if err != nil {
				return
			}
			mediaForm := responseForm.MediasForm{
				MediaURL:  media.MediaURL,
				MediaType: model.GetMediaTypeByInt(media.MediaType),
			}
			mediasForms = append(mediasForms, mediaForm)
		}
	}
	form = responseForm.GetMomentResponseForm{
		NickName:   user.Nickname,
		Avatar:     user.Avatar,
		UserID:     user.UserID.String(),
		Signature:  user.UserInfo.Signature,
		MomentID:   moment.MomentID,
		MomentInfo: moment.MomentInfo,
		Location:   moment.Location,
		ReadNum:    moment.ReadNum,
		LikeNum:    0,
		Medias:     mediasForms,
		CreatedAt:  moment.CreatedAt.String(),
	}
	return
}
func (m *MomentStruct) ExistLike() bool {
	return model.ExistLikeMap(m.MomentID, m.UserID, model.LikeMoment)
}
func (m *MomentStruct) Like() (err error) {
	_, err = model.AddLikeMap(m.MomentID, m.UserID, model.LikeMoment)
	return
}
