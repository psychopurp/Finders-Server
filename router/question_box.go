package router

import (
	v1 "finders-server/api/v1"
	"finders-server/middleware"
	"github.com/gin-gonic/gin"
)

func InitQuestionBoxRouter(Router *gin.RouterGroup) {

	QuestionBoxRouter := Router.Group("question_box")
	{
		QuestionBoxRouter.POST("create_question_box", middleware.JWT(), v1.CreateQuestionBox)
		QuestionBoxRouter.POST("like_question_box", middleware.JWT(), v1.LikeQuestionBox)
	}
}

