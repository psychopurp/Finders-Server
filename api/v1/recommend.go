package v1

import (
	"finders-server/global/response"
	"finders-server/model"
	"finders-server/model/responseForm"
	"github.com/gin-gonic/gin"
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
