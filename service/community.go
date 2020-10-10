package service

import (
	"finders-server/model"
)

type CommunityStruct struct {
	CommunityID          int
	CommunityCreator     string
	CommunityName        string
	CommunityDescription string
	CommunityStatus      int
	Background           string
	CommunityAvatar      string

	PageNum  int
	PageSize int
	Page     int
}

func (communityStruct *CommunityStruct) Exist() bool {
	maps := map[string]interface{}{
		"community_creator":     communityStruct.CommunityCreator,
		"community_name":        communityStruct.CommunityName,
		"community_description": communityStruct.CommunityDescription,
		"background":            communityStruct.Background,
	}
	return model.ExistCommunityByMap(maps)
}

func (communityStruct *CommunityStruct) Add() (com model.Community, err error) {

	community := &model.Community{
		CommunityCreator:     communityStruct.CommunityCreator,
		CommunityName:        communityStruct.CommunityName,
		CommunityDescription: communityStruct.CommunityDescription,
		Background:           communityStruct.Background,
		CommunityAvatar:      communityStruct.CommunityAvatar,
	}
	return model.AddCommunityByMap(community)
}

func (communityStruct *CommunityStruct) ExistByID() (isExist bool, err error) {
	isExist, err = model.ExistCommunityByID(communityStruct.CommunityID)
	return
}

func (communityStruct *CommunityStruct) Edit() (err error) {
	// 需要使用 结构体修改数据 否则会修改所有字段 用结构体修改则会自动省去空字段
	community := model.Community{
		CommunityName:        communityStruct.CommunityName,
		CommunityDescription: communityStruct.CommunityDescription,
		Background:           communityStruct.Background,
		CommunityAvatar:      communityStruct.CommunityAvatar,
	}
	return model.UpdateCommunityByCommunity(communityStruct.CommunityID, community)
}
