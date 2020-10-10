package service

import (
	"finders-server/model"
	"strconv"
)

type TagStruct struct {
	TagID   int
	TagName string
	TagType int
	Base
}

func (t *TagStruct) AddQuestionBoxTagByName(tagNames []string) (tagIDs []int, err error) {
	var (
		tag *model.Tag
	)
	for _, tagName := range tagNames {
		tag, err = t.Affair.FirstTagOrCreate(tagName, model.TagQuestionBoxDIY)
		if err != nil {
			return
		}
		tagIDs = append(tagIDs, tag.TagID)
	}
	return
}

func (t *TagStruct) AddTagMap(tagID, itemType int, itemID string) (err error) {
	tagMap := &model.TagMap{
		ItemID:   itemID,
		ItemType: itemType,
		TagID:    tagID,
	}
	return t.Affair.AddTagMap(tagMap)
}

func (t *TagStruct) AddQuestionBoxTagMap(questionBoxID int, tagIDs []int) (err error) {
	itemID := strconv.Itoa(questionBoxID)
	for _, tagID := range tagIDs {
		err = t.AddTagMap(tagID, model.TagQuestionBoxType, itemID)
		if err != nil {
			return
		}
	}
	return
}
