package communityService

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
	maps := map[string]interface{}{
		"community_creator":     communityStruct.CommunityCreator,
		"community_name":        communityStruct.CommunityName,
		"community_description": communityStruct.CommunityDescription,
		"background":            communityStruct.Background,
	}
	return model.AddCommunityByMap(maps)
}

func (communityStruct *CommunityStruct) ExistByID() (isExist bool, err error) {
	isExist, err = model.ExistCommunityByID(communityStruct.CommunityID)
	return
}

func (communityStruct *CommunityStruct) Edit() (err error) {
	community := model.Community{
		CommunityName:        communityStruct.CommunityName,
		CommunityDescription: communityStruct.CommunityDescription,
		Background:           communityStruct.Background,
	}
	return model.UpdateCommunityByCommunity(communityStruct.CommunityID, community)
}
