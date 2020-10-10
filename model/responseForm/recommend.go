package responseForm

type MainRecommendResponseForm struct {
	Cnt   int          `json:"cnt"`
	Cards []SimpleCard `json:"cards"`
}

type SimpleCard struct {
	CardID   int    `json:"card_id"`
	ItemID   string `json:"item_id"`
	ItemType int    `json:"item_type"`
}

type UserInfoCard struct {
	UserID            string           `json:"user_id"`
	Avatar            string           `json:"avatar"`
	NickName          string           `json:"nick_name"`
	Signature         string           `json:"signature"`
	SharedCommunities []ShareCommunity `json:"shared_communities"`
}

type ShareCommunity struct {
	CommunityID   int    `json:"community_id"`
	CommunityName string `json:"community_name"`
}

type ActivityCard struct {
	ActivityTitle string       `json:"activity_title"`
	ActivityInfo  string       `json:"activity_info"`
	NickName      string       `json:"nick_name"`
	UserID        string       `json:"user_id"`
	Avatar        string       `json:"avatar"`
	CommunityID   int          `json:"community_id"`
	CommunityName string       `json:"community_name"`
	Medias        []MediasForm `json:"medias"`
}


type MomentCard struct {
	NickName   string       `json:"nick_name"`
	Avatar     string       `json:"avatar"`
	UserID     string       `json:"user_id"`
	Signature  string       `json:"signature"`
	MomentID   string       `json:"moment_id"`
	MomentInfo string       `json:"moment_info"`
	Location   string       `json:"location"`
	Medias     []MediasForm `json:"medias"`
	CreatedAt  string       `json:"created_at"`
}