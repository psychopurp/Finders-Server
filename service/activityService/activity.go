package activityService

import (
	"finders-server/model"
	"finders-server/model/responseForm"
	"finders-server/pkg/e"
	"finders-server/utils"
	"math"
	"strings"
)

type ActivityStruct struct {
	ActivityID     string
	ActivityStatus int
	ActivityInfo   string
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
		"activity_info": activityStruct.ActivityInfo,
		"media_id":      activityStruct.MediaIDs,
		"user_id":       activityStruct.UserID,
		"community_id":  activityStruct.CommunityID,
	}
	activity, err = model.AddActivityByMap(data)
	return
}

func (activityStruct *ActivityStruct) GetAll() (activities []*model.Activity, err error) {
	activities, err = model.GetActivitiesByCommunityID(activityStruct.PageNum, activityStruct.PageSize, activityStruct.CommunityID)
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
	form = responseForm.ActivityInfoForm{
		ActivityID:   activity.ActivityID,
		ActivityInfo: activity.ActivityInfo,
		CollectNum:   activity.CollectNum,
		CommentNum:   activity.CommentNum,
		ReadNum:      activity.ReadNum,
		Tags:         TagInfoForms,
		Medias:       mediasForms,
		NickName:     user.Nickname,
		UserID:       user.UserID.String(),
		Avatar:       user.Avatar,
		UserType:     userType,
		CreatedAt:    activity.CreatedAt.String(),
	}
	return
}

func (activityStruct *ActivityStruct) GetActivitiesPageResponse() (form responseForm.ActivitiesResponseForm, err error) {
	var (
		activities       []*model.Activity
		totalActivityCNT int
		activitiesForms  []responseForm.ActivityInfoForm

		ok bool
	)
	form.Page = activityStruct.Page
	totalActivityCNT, _ = activityStruct.CountByCommunityID()
	form.TotalCNT = totalActivityCNT
	if totalActivityCNT == 0 {
		return
	}
	form.TotalPage = int(math.Ceil(float64(totalActivityCNT) / float64(activityStruct.PageSize)))
	activities, err = activityStruct.GetAll()
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
		var activitiesForm = responseForm.ActivityInfoForm{
			ActivityID:   activity.ActivityID,
			ActivityInfo: activity.ActivityInfo,
			CollectNum:   activity.CollectNum,
			CommentNum:   activity.CommentNum,
			ReadNum:      activity.ReadNum,
			Tags:         TagInfoForms,
			Medias:       mediasForms,
			NickName:     activity.User.Nickname,
			UserID:       activity.User.UserID.String(),
			Avatar:       activity.User.Avatar,
			UserType:     userType,
			CreatedAt:    activity.CreatedAt.String(),
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

func (activityStruct *ActivityStruct) ExistLike() bool {
	return model.ExistActivityLike(activityStruct.ActivityID, activityStruct.UserID)
}

func (activityStruct *ActivityStruct) Like() (like model.ActivityLike, err error) {
	like, err = model.AddActivityLike(activityStruct.ActivityID, activityStruct.UserID)
	return
}

func (activityStruct *ActivityStruct) DisLike() (err error) {
	err = model.DeleteActivityLike(activityStruct.ActivityID, activityStruct.UserID)
	return
}

func (activityStruct *ActivityStruct) CountLike() (cnt int, err error) {
	cnt, err = model.GetActivityLikeTotal(activityStruct.UserID)
	return
}

func (activityStruct *ActivityStruct) GetALlActivityLikes() (activityLikes []*model.ActivityLike, err error) {
	activityLikes, err = model.GetActivityLikesByUserID(activityStruct.UserID)
	return
}

func (activityStruct *ActivityStruct) GetAllLikeActivities() (activities []*model.Activity, err error) {
	var (
		activityLikes []*model.ActivityLike
		activityIDs   []string
	)
	activityLikes, err = activityStruct.GetALlActivityLikes()
	if err != nil {
		return
	}
	for _, activityLike := range activityLikes {
		activityIDs = append(activityIDs, activityLike.ActivityID)
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
