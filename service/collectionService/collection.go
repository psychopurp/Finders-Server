package collectionService

import (
	"finders-server/model"
	"finders-server/model/responseForm"
	"finders-server/pkg/e"
	"finders-server/utils"
	"fmt"
	"math"
)

type CollectionStruct struct {
	CollectionID     int
	CollectionStatus string
	CollectionType   int
	UserID           string
	Link             string

	PageNum  int
	PageSize int
	Page     int
}

func (collectionStruct *CollectionStruct) AddCollectNum() (err error) {
	err = model.AddActivityCollectNum(collectionStruct.Link, model.AddOP)
	return
}
func (collectionStruct *CollectionStruct) CutCollectNum() (err error) {
	err = model.AddActivityCollectNum(collectionStruct.Link, model.MinusOP)
	return
}

func (collectionStruct *CollectionStruct) Exist() bool {
	data := map[string]interface{}{
		"collection_type": collectionStruct.CollectionType,
		"user_id":         collectionStruct.UserID,
		"link":            collectionStruct.Link,
	}
	return model.ExistCollectionByMap(data)
}
func (collectionStruct *CollectionStruct) AddCollection() (collection model.Collection, err error) {
	data := map[string]interface{}{
		"collection_type": collectionStruct.CollectionType,
		"user_id":         collectionStruct.UserID,
		"link":            collectionStruct.Link,
	}
	collection, err = model.AddCollectionByMap(data)
	return
}

func (collectionStruct *CollectionStruct) RemoveCollection() (err error) {
	data := map[string]interface{}{
		"collection_type": collectionStruct.CollectionType,
		"user_id":         collectionStruct.UserID,
		"link":            collectionStruct.Link,
	}
	err = model.DeleteCollectionByMap(data)
	return
}

func (collectionStruct *CollectionStruct) GetAllCommunities() (communities []*model.Community, err error) {
	var collections []*model.Collection
	var links []string
	collections, err = model.GetCollectionIDs(collectionStruct.UserID, model.CollectionCommunity)
	if err != nil {
		return
	}
	for _, collection := range collections {
		links = append(links, collection.Link)
	}
	communities, err = model.GetCommunities(collectionStruct.PageNum, collectionStruct.PageSize, links)
	return
}

func (collectionStruct *CollectionStruct) GetAllActivities() (activities []*model.Activity, err error) {
	var collections []*model.Collection
	var links []string
	collections, err = model.GetCollectionIDs(collectionStruct.UserID, model.CollectionActivity)
	if err != nil {
		return
	}
	for _, collection := range collections {
		links = append(links, collection.Link)
	}
	activities, err = model.GetActivitiesByActivityIDs(collectionStruct.PageNum, collectionStruct.PageSize, links)
	return
}

func (collectionStruct *CollectionStruct) GetCommunitiesCollectionResponse() (form responseForm.CommunitiesResponseForm, err error) {
	var (
		communities       []*model.Community
		totalCommunityCNT int
		communitiesForms  []responseForm.CommunitiesForm
	)
	// 设置返回第几页
	form.Page = collectionStruct.Page
	// 获取所有收藏
	totalCommunityCNT, _ = model.GetCollectionTotal(collectionStruct.UserID, model.CollectionCommunity)
	form.TotalCNT = totalCommunityCNT
	// 若没有数据直接返回
	if totalCommunityCNT == 0 {
		return
	}
	// 计算总页数
	form.TotalPage = int(math.Ceil(float64(totalCommunityCNT) / float64(collectionStruct.PageSize)))
	// 根据分页获取一定数量的community
	communities, err = collectionStruct.GetAllCommunities()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "GetCommunitiesCollectionResponse 1")
		return
	}
	// 返回的个数
	form.CNT = len(communities)
	for _, community := range communities {
		var user model.User
		user, err = model.GetUserByUserID(community.CommunityCreator)
		if err != nil {
			return
		}
		var communitiesForm = responseForm.CommunitiesForm{
			CommunityID:          community.CommunityID,
			CommunityCreator:     community.CommunityCreator,
			NickName:             user.Nickname,
			Avatar:               user.Avatar,
			CommunityName:        community.CommunityName,
			CommunityDescription: community.CommunityDescription,
			Background:           community.Background,
		}
		communitiesForms = append(communitiesForms, communitiesForm)
	}
	form.CommunitiesForms = communitiesForms
	return
}

func (collectionStruct *CollectionStruct) GetActivityCollectionResponse() (form responseForm.ActivitiesResponseForm, err error) {
	var (
		activities       []*model.Activity
		totalActivityCNT int
		activitiesForms  []responseForm.ActivityInfoForm
		ok               bool
	)
	form.Page = collectionStruct.Page
	totalActivityCNT, _ = model.GetCollectionTotal(collectionStruct.UserID, model.CollectionActivity)
	form.TotalCNT = totalActivityCNT
	if totalActivityCNT == 0 {
		return
	}
	form.TotalPage = int(math.Ceil(float64(totalActivityCNT) / float64(collectionStruct.PageSize)))
	activities, err = collectionStruct.GetAllActivities()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "GetActivitiesPageResponse1")
		return
	}
	form.CNT = len(activities)
	var getMediaType = map[int]string{
		model.PICTURE: "picture",
		model.VIDEO:   "video",
	}
	for _, activity := range activities {
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
		userType := "normal"
		ok, err = model.IsManagerByUserID(activity.UserID)
		if err != nil {
			return
		}
		if ok {
			userType = "manager"
		}
		fmt.Println(activity.User.Nickname)
		var activitiesForm = responseForm.ActivityInfoForm{
			ActivityID:   activity.ActivityID,
			ActivityInfo: activity.ActivityInfo,
			CollectNum:   activity.CollectNum,
			CommentNum:   activity.CommentNum,
			ReadNum:      activity.ReadNum,
			Tags:         TagInfoForms,
			MediaURL:     activity.Media.MediaURL,
			MediaType:    getMediaType[activity.Media.MediaType],
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
