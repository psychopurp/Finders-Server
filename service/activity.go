package service

import (
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/model/responseForm"
	"finders-server/pkg/e"
	"finders-server/st"
	"finders-server/utils"
	"github.com/gin-gonic/gin"
	"math"
	"strings"
)

type ActivityStruct struct {
	ActivityID     string
	ActivityStatus int
	ActivityInfo   string
	ActivityTitle  string
	CollectNum     int
	CommentNum     int
	ReadNum        int
	Media          model.Media
	MediaIDs       string
	UserID         string
	CommunityID    int

	PageNum  int
	PageSize int
	Page     int

	Affair *model.AffairService
}

func (activityStruct *ActivityStruct) AffairInit(c *gin.Context) (err error) {
	activityStruct.Affair = new(model.AffairService)
	err = activityStruct.Affair.NewAffairs()
	if err != nil {
		response.FailWithMsg(e.MYSQL_ERROR, c)
		return
	}
	return nil
}

func (activityStruct *ActivityStruct) AffairBegin() func() {
	return activityStruct.Affair.DeferFunc()
}

func (activityStruct *ActivityStruct) AddReadNum() (err error) {
	err = model.UpdateActivityReadNum(activityStruct.ActivityID, model.AddOP)
	return
}
func (activityStruct *ActivityStruct) CutReadNum() (err error) {
	err = model.UpdateActivityReadNum(activityStruct.ActivityID, model.MinusOP)
	return
}

func (activityStruct *ActivityStruct) Exist() bool {
	data := map[string]interface{}{
		"activity_info": activityStruct.ActivityInfo,
		"media_id":      activityStruct.MediaIDs,
		"user_id":       activityStruct.UserID,
		"community_id":  activityStruct.CommunityID,
	}
	return model.ExistActivityByMap(data)
}

func (activityStruct *ActivityStruct) ExistByID() bool {
	data := map[string]interface{}{
		"activity_id": activityStruct.ActivityID,
	}
	return model.ExistActivityByMap(data)
}

func (activityStruct *ActivityStruct) Add() (activity model.Activity, err error) {
	data := map[string]interface{}{
		"activity_info":  activityStruct.ActivityInfo,
		"activity_title": activityStruct.ActivityTitle,
		"media_id":       activityStruct.MediaIDs,
		"user_id":        activityStruct.UserID,
		"community_id":   activityStruct.CommunityID,
	}
	activity, err = model.AddActivityByMap(data)
	return
}

func (activityStruct *ActivityStruct) GetActivityLikeNum() (cnt int, err error) {
	return model.GetLikeMapNumByObjectID(activityStruct.ActivityID)
}

func (activityStruct *ActivityStruct) IsLikeActivity() bool {
	return model.ExistLikeMap(activityStruct.ActivityID, activityStruct.UserID, model.LikeActivity)
}

func (activityStruct *ActivityStruct) GetAllByCommunityID() (activities []*model.Activity, err error) {
	activities, err = model.GetActivitiesByCommunityID(activityStruct.PageNum, activityStruct.PageSize, activityStruct.CommunityID)
	return
}

func (activityStruct *ActivityStruct) GetAllByUserID() (activities []*model.Activity, err error) {
	activities, err = model.GetActivitiesByUserID(activityStruct.PageNum, activityStruct.PageSize, activityStruct.UserID)
	return
}

func (activityStruct *ActivityStruct) GetByID() (activity model.Activity, err error) {
	activity, err = model.GetActivityByID(activityStruct.ActivityID)
	return
}

func (activityStruct *ActivityStruct) GetActivityInfoResponse() (form responseForm.ActivityInfoForm, err error) {
	var (
		activity     model.Activity
		ok           bool
		TagInfoForms []responseForm.TagInfoForm
		mediasForms  []responseForm.MediasForm
	)
	activity, err = activityStruct.GetByID()
	if err != nil {
		return
	}

	var tags []*model.Tag
	var user model.User
	tags, err = model.GetTagsByActivityID(activity.ActivityID)
	if err != nil {
		st.DebugWithFuncName(err)
		return
	}
	user, err = model.GetUserByUserID(activity.UserID)
	if err != nil {
		return
	}
	for _, tag := range tags {
		tagInfoForm := responseForm.TagInfoForm{
			TagName: tag.TagName,
			TagType: tag.TagType,
		}
		TagInfoForms = append(TagInfoForms, tagInfoForm)
	}
	userType := model.CommunityNormalMember
	ok, err = model.IsManagerByUserID(activity.UserID)
	if err != nil {
		return
	}
	if ok {
		userType = model.CommunityManagerMember
	}
	if len(activity.MediaIDs) != 0 {
		mediaIDs := strings.Split(activity.MediaIDs, ";")
		for _, mediaID := range mediaIDs {
			if mediaID == "" {
				continue
			}
			var media model.Media
			media, err = model.GetMediaByMediaID(mediaID)
			if err != nil {
				st.DebugWithFuncName(err)
				return
			}
			mediaForm := responseForm.MediasForm{
				MediaURL:  media.MediaURL,
				MediaType: model.GetMediaTypeByInt(media.MediaType),
			}
			mediasForms = append(mediasForms, mediaForm)
		}
	}
	var community model.Community
	community, err = model.GetCommunityByCommunityID(activity.CommunityID)
	if err != nil {
		return
	}
	form = responseForm.ActivityInfoForm{
		ActivityID:    activity.ActivityID,
		ActivityInfo:  activity.ActivityInfo,
		ActivityTitle: activity.ActivityTitle,
		CollectNum:    activity.CollectNum,
		CommentNum:    activity.CommentNum,
		ReadNum:       activity.ReadNum,
		Tags:          TagInfoForms,
		Medias:        mediasForms,
		NickName:      user.Nickname,
		UserID:        user.UserID.String(),
		Avatar:        user.Avatar,
		Signature:     user.UserInfo.Signature,
		UserType:      userType,
		CreatedAt:     activity.CreatedAt.String(),

		CommunityID:          community.CommunityID,
		CommunityName:        community.CommunityName,
		Background:           community.Background,
		CommunityDescription: community.CommunityDescription,
	}
	return
}

const (
	GetActivitiesOnCommunity = 1 + iota
	GetActivitiesOnUser      = 1 + iota
)

func (activityStruct *ActivityStruct) GetActivitiesPageResponse(filterType int) (form responseForm.ActivitiesResponseForm, err error) {
	var (
		activities       []*model.Activity
		totalActivityCNT int
		activitiesForms  []responseForm.ActivityInfoForm
		ok               bool
	)
	form.Page = activityStruct.Page
	if filterType == GetActivitiesOnCommunity {
		totalActivityCNT, _ = activityStruct.CountByCommunityID()
	} else if filterType == GetActivitiesOnUser {
		totalActivityCNT, _ = activityStruct.CountByUserID()
	}
	form.TotalCNT = totalActivityCNT
	if totalActivityCNT == 0 {
		return
	}
	form.TotalPage = int(math.Ceil(float64(totalActivityCNT) / float64(activityStruct.PageSize)))
	if filterType == GetActivitiesOnCommunity {
		activities, err = activityStruct.GetAllByCommunityID()
	} else if filterType == GetActivitiesOnUser {
		activities, err = activityStruct.GetAllByUserID()
	}
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "GetActivitiesPageResponse1")
		return
	}
	form.CNT = len(activities)

	for _, activity := range activities {
		var mediasForms []responseForm.MediasForm
		var tags []*model.Tag
		tags, err = model.GetTagsByActivityID(activity.ActivityID)
		if err != nil {
			return
		}
		var TagInfoForms []responseForm.TagInfoForm
		for _, tag := range tags {
			tagInfoForm := responseForm.TagInfoForm{
				TagName: tag.TagName,
				TagType: tag.TagType,
			}
			TagInfoForms = append(TagInfoForms, tagInfoForm)
		}
		userType := model.CommunityNormalMember
		ok, err = model.IsManagerByUserID(activity.UserID)
		if err != nil {
			return
		}
		if ok {
			userType = model.CommunityManagerMember
		}
		if len(activity.MediaIDs) != 0 {
			mediaIDs := strings.Split(activity.MediaIDs, ";")
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
		var community model.Community
		community, err = model.GetCommunityByCommunityID(activity.CommunityID)
		if err != nil {
			return
		}
		var activitiesForm = responseForm.ActivityInfoForm{
			ActivityID:    activity.ActivityID,
			ActivityInfo:  activity.ActivityInfo,
			CollectNum:    activity.CollectNum,
			CommentNum:    activity.CommentNum,
			ReadNum:       activity.ReadNum,
			Tags:          TagInfoForms,
			Medias:        mediasForms,
			NickName:      activity.User.Nickname,
			UserID:        activity.User.UserID.String(),
			Avatar:        activity.User.Avatar,
			UserType:      userType,
			CreatedAt:     activity.CreatedAt.String(),
			CommunityID:   community.CommunityID,
			CommunityName: community.CommunityName,
			Background:    community.Background,
		}
		activitiesForms = append(activitiesForms, activitiesForm)
	}
	form.ActivitiesForms = activitiesForms
	return
}

func (activityStruct *ActivityStruct) CountByCommunityID() (cnt int, err error) {
	cnt, err = model.GetActivityTotalByCommunityID(activityStruct.CommunityID)
	return
}

func (activityStruct *ActivityStruct) CountByUserID() (cnt int, err error) {
	cnt, err = model.GetActivityTotalByUserID(activityStruct.UserID)
	return
}

func (activityStruct *ActivityStruct) ExistLike() bool {
	return model.ExistLikeMap(activityStruct.ActivityID, activityStruct.UserID, model.LikeActivity)
}

func (activityStruct *ActivityStruct) Like() (like model.LikeMap, err error) {
	like, err = model.AddLikeMap(activityStruct.ActivityID, activityStruct.UserID, model.LikeActivity)
	return
}

func (activityStruct *ActivityStruct) DisLike() (err error) {
	err = model.DeleteLikeMap(activityStruct.ActivityID, activityStruct.UserID, model.LikeActivity)
	return
}

func (activityStruct *ActivityStruct) CountLike() (cnt int, err error) {
	cnt, err = model.GetUserLikeMapTotal(activityStruct.UserID, model.LikeActivity)
	return
}

func (activityStruct *ActivityStruct) GetALlActivityLikes() (likeMaps []*model.LikeMap, err error) {
	likeMaps, err = model.GetLikeMapsByUserID(activityStruct.UserID, model.LikeActivity)
	return
}

func (activityStruct *ActivityStruct) GetAllLikeActivities() (activities []*model.Activity, err error) {
	var (
		likeMaps    []*model.LikeMap
		activityIDs []string
	)
	likeMaps, err = activityStruct.GetALlActivityLikes()
	if err != nil {
		return
	}
	for _, activityLike := range likeMaps {
		activityIDs = append(activityIDs, activityLike.ObjectID)
	}
	activities, err = model.GetActivitiesByActivityIDs(activityStruct.PageNum, activityStruct.PageSize, activityIDs)
	return
}
func (activityStruct *ActivityStruct) GetActivityLikesResponse() (form responseForm.ActivityLikesResponseForm, err error) {
	var (
		totalLikeCNT int
		activities   []*model.Activity
		likeForms    []responseForm.LikeForm
	)
	form.Page = activityStruct.Page
	totalLikeCNT, _ = activityStruct.CountLike()
	form.TotalCNT = totalLikeCNT
	if totalLikeCNT == 0 {
		return
	}
	form.TotalPage = int(math.Ceil(float64(totalLikeCNT) / float64(activityStruct.PageSize)))
	activities, err = activityStruct.GetAllLikeActivities()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "GetActivityLikesResponse1")
		return
	}
	form.CNT = len(activities)
	for _, activity := range activities {
		var likeForm = responseForm.LikeForm{
			NickName: activity.User.Nickname,
			UserID:   activity.User.UserID.String(),
			Avatar:   activity.User.Avatar,
		}
		likeForms = append(likeForms, likeForm)
	}
	form.LikeForms = likeForms
	return
}