package commentService

import (
	"finders-server/model"
	"finders-server/model/responseForm"
	"finders-server/pkg/e"
	"finders-server/utils"
	"github.com/unknwon/com"
	"math"
	"strconv"
)

type CommentStruct struct {
	CommentID int
	ItemID    string
	ItemType  int
	Content   string
	ToUID     string
	FromUID   string
	Status    int

	PageNum  int
	PageSize int
	Page     int
}

func (commentStruct *CommentStruct) AddCommentNum() (err error) {
	err = model.UpdateActivityCommentNum(commentStruct.ItemID, model.AddOP)
	return
}

func (commentStruct *CommentStruct) CutCommentNum() (err error) {
	err = model.UpdateActivityCommentNum(commentStruct.ItemID, model.MinusOP)
	return
}

func (commentStruct *CommentStruct) Exist() bool {
	data := map[string]interface{}{
		"item_id":  commentStruct.ItemID,
		"content":  commentStruct.Content,
		"from_uid": commentStruct.FromUID,
	}
	return model.ExistCommentByMap(data)
}

func (commentStruct *CommentStruct) AddComment() (comment model.Comment, err error) {
	var activity model.Activity
	activity, err = model.GetActivityByID(commentStruct.ItemID)
	if err != nil {
		return
	}
	data := map[string]interface{}{
		"item_id":   commentStruct.ItemID,
		"content":   commentStruct.Content,
		"item_type": model.CommentOnActivity,
		"from_uid":  commentStruct.FromUID,
		"to_uid":    activity.UserID,
	}
	comment, err = model.AddCommentByMap(data)
	return
}

func (commentStruct *CommentStruct) AddReply() (comment model.Comment, err error) {
	var replyComment model.Comment
	replyComment, err = model.GetCommentByCommentID(com.StrTo(commentStruct.ItemID).MustInt())
	if err != nil {
		return
	}
	data := map[string]interface{}{
		"item_id":   commentStruct.ItemID,
		"content":   commentStruct.Content,
		"item_type": model.CommentOnComment,
		"from_uid":  commentStruct.FromUID,
		"to_uid":    replyComment.FromUID,
	}
	comment, err = model.AddCommentByMap(data)
	return
}
func (commentStruct *CommentStruct) CountComment() (cnt int, err error) {
	return model.GetCommentTotal(commentStruct.ItemID, commentStruct.ItemType)
}

func (commentStruct *CommentStruct) GetAllComment() (comments []*model.Comment, err error) {
	return model.GetCommentsByItemID(commentStruct.PageNum, commentStruct.PageSize, commentStruct.ItemID, commentStruct.ItemType)
}

func (commentStruct *CommentStruct) GetCommentsResponse() (form responseForm.CommentResponseForm, err error) {
	var (
		comments     []*model.Comment
		totalCNT     int
		commentForms []responseForm.CommentInfoForm
	)
	form.Page = commentStruct.Page
	totalCNT, err = commentStruct.CountComment()
	if err != nil {
		return
	}
	form.TotalCNT = totalCNT
	if totalCNT == 0 {
		return
	}
	form.TotalPage = int(math.Ceil(float64(totalCNT) / float64(commentStruct.PageSize)))
	comments, err = commentStruct.GetAllComment()
	if err != nil {
		err = utils.GetErrorAndLog(e.MYSQL_ERROR, err, "GetActivityCommentResponse1")
		return
	}
	form.CNT = len(comments)
	for _, comment := range comments {
		var replyNum int
		if commentStruct.ItemType == model.CommentOnActivity {
			replyNum, err = model.GetCommentTotal(strconv.Itoa(comment.CommentID), model.CommentOnComment)
			if err != nil {
				return
			}
		}
		var commentForm = responseForm.CommentInfoForm{
			CommentID: comment.CommentID,
			Content:   comment.Content,
			NickName:  comment.FromUser.Nickname,
			UserID:    comment.FromUser.UserID.String(),
			Avatar:    comment.FromUser.Avatar,
			CreatedAt: comment.CreatedAt.String(),
			ReplyNum:  replyNum,
		}
		commentForms = append(commentForms, commentForm)
	}
	form.CommentForms = commentForms
	return
}
