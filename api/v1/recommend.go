package v1

import (
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/model/responseForm"
	"finders-server/pkg/e"
	"finders-server/service"
	"finders-server/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

func MainRecommend(c *gin.Context) {
	var (
		users      []*model.User
		activities []*model.Activity
		moments    []*model.Moment
		form       responseForm.MainRecommendResponseForm
	)
	rand.Seed(time.Now().Unix())
	users, _ = model.GetUsers()
	activities, _ = model.GetActivities()
	moments, _ = model.GetMoments()
	cnt := 1
	const (
		U = iota + 1
		A
		M
	)
	//
	for i := 0; i < 2; i++ {
		card := responseForm.SimpleCard{
			CardID:   cnt,
			ItemID:   users[rand.Intn(len(users))].UserID.String(),
			ItemType: U,
		}
		cnt++
		form.Cards = append(form.Cards, card)
	}
	for i := 0; i < 3; i++ {
		card := responseForm.SimpleCard{
			CardID:   cnt,
			ItemID:   activities[rand.Intn(len(activities))].ActivityID,
			ItemType: A,
		}
		cnt++
		form.Cards = append(form.Cards, card)
	}
	for i := 0; i < 3; i++ {
		card := responseForm.SimpleCard{
			CardID:   cnt,
			ItemID:   moments[rand.Intn(len(moments))].MomentID,
			ItemType: M,
		}
		cnt++
		form.Cards = append(form.Cards, card)
	}
	form.Cnt = 8
	sol := &Solution{}
	sol.nums = form.Cards
	form.Cards = sol.Shuffle()
	response.OkWithData(form, c)
}

type Solution struct {
	nums []responseForm.SimpleCard
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []responseForm.SimpleCard {
	nums := make([]responseForm.SimpleCard, len(this.nums))
	copy(nums, this.nums)
	rand.Shuffle(len(nums), func(i int, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

func UserRecommend(c *gin.Context) {
	var (
		err    error
		form   responseForm.UserInfoCard
		userID string
		user   model.User
	)
	userID = c.Query("userId")
	if userID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	userStruct := service.UserStruct{
		UserID: uuid.FromStringOrNil(userID),
	}
	user, err = userStruct.GetUserByUserID()
	if utils.FailOnError(e.MYSQL_ERROR, err, c) {
		return
	}
	form = responseForm.UserInfoCard{
		UserID:    userID,
		Avatar:    user.Avatar,
		NickName:  user.Nickname,
		Signature: user.UserInfo.Signature,
		SharedCommunities: []responseForm.ShareCommunity{
			{
				CommunityID:   10,
				CommunityName: "尴尬",
			},
			{
				CommunityID:   11,
				CommunityName: "郁闷",
			},
			{
				CommunityID:   12,
				CommunityName: "佛",
			},
		},
	}
	response.OkWithData(form, c)
}

func ActivityRecommend(c *gin.Context) {
	var (
		err        error
		form       responseForm.ActivityCard
		tmp        responseForm.ActivityInfoForm
		ActivityID string
	)
	ActivityID = c.Query("activity_id")
	if ActivityID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	activityStruct := service.ActivityStruct{ActivityID: ActivityID}
	tmp, err = activityStruct.GetActivityInfoResponse()
	if utils.FailOnError(e.MYSQL_ERROR, err, c) {
		return
	}
	form = responseForm.ActivityCard{
		ActivityTitle: tmp.ActivityTitle,
		ActivityInfo:  tmp.ActivityInfo,
		NickName:      tmp.NickName,
		UserID:        tmp.UserID,
		Avatar:        tmp.Avatar,
		CommunityID:   tmp.CommunityID,
		CommunityName: tmp.CommunityName,
		Medias:        tmp.Medias,
	}
	response.OkWithData(form, c)
}

func MomentRecommend(c *gin.Context) {
	var (
		err      error
		form     responseForm.MomentCard
		tmp      responseForm.GetMomentResponseForm
		momentID string
	)
	momentID = c.Query("moment_id")
	if momentID == "" {
		response.FailWithMsg(e.INFO_ERROR, c)
		return
	}
	momentStruct := service.MomentStruct{MomentID: momentID}
	tmp, err = momentStruct.GetMomentResponseForm()
	if utils.FailOnError(e.MYSQL_ERROR, err, c) {
		return
	}
	form = responseForm.MomentCard{
		NickName:   tmp.NickName,
		Avatar:     tmp.Avatar,
		UserID:     tmp.UserID,
		Signature:  tmp.Signature,
		MomentID:   tmp.MomentID,
		MomentInfo: tmp.MomentInfo,
		Location:   tmp.Location,
		Medias:     tmp.Medias,
		CreatedAt:  tmp.CreatedAt,
	}
	response.OkWithData(form, c)
}
