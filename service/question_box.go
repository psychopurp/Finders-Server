package service

import (
	"finders-server/model"
	"finders-server/pkg/e"
	"fmt"
	"strconv"
)

type QuestionBoxStruct struct {
	QuestionBoxID     int
	UserID            string
	QuestionBoxStatus int64
	QuestionBoxInfo   string
	UseNum            int
	ReplyNum          int
	LikeNum           int
	TagNames          string
	Base
}

func (q *QuestionBoxStruct) Add() (err error) {
	questionBox := &model.QuestionBox{
		UserID:            q.UserID,
		QuestionBoxStatus: model.QuestionStatusNormal,
		QuestionBoxInfo:   q.QuestionBoxInfo,
		UseNum:            0,
		ReplyNum:          0,
		LikeNum:           0,
		TagNames:          q.TagNames,
	}
	err = q.Affair.AddQuestionBox(questionBox)
	if err != nil {
		return
	}
	q.QuestionBoxID = questionBox.QuestionBoxID
	return
}

func (q *QuestionBoxStruct) CheckExistQuestionInfo() bool {
	return model.ExistQuestionBoxByInfo(q.QuestionBoxInfo)
}


func (q *QuestionBoxStruct)AddLike()(err error){
	if model.ExistLikeMap(strconv.Itoa(q.QuestionBoxID), q.UserID, model.LikeQuestionBox){
		return fmt.Errorf(e.REPEAT_SUBMIT)
	}
	return nil
}