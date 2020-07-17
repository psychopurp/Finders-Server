package responseForm

type CommunitiesResponseForm struct {
	Paginator
	CommunitiesForms []CommunitiesForm ` json:"communities"`
}

type CommunitiesForm struct {
	CommunityID          int    `json:"community_id"`
	CommunityCreator     string `json:"community_creator"`
	NickName             string `json:"nick_name"`
	Avatar               string `json:"avatar"`
	CommunityName        string `json:"community_name"`
	CommunityDescription string `json:"community_description"`
	Background           string `json:"backgroud"`
}
