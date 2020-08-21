package responseForm

type ActivitiesResponseForm struct {
	Paginator
	ActivitiesForms []ActivityInfoForm `json:"activities"`
}

type ActivityInfoForm struct {
	ActivityID    string        `json:"activity_id"`
	ActivityInfo  string        `json:"activity_info"`
	ActivityTitle string        `json:"activity_title"`
	CollectNum    int           `json:"collect_num"`
	CommentNum    int           `json:"comment_num"`
	ReadNum       int           `json:"read_num"`
	Tags          []TagInfoForm `json:"tags"`
	Medias        []MediasForm  `json:"medias"`
	NickName      string        `json:"nick_name"`
	UserID        string        `json:"user_id"`
	Avatar        string        `json:"avatar"`
	UserType      string        `json:"user_type"`
	CreatedAt     string        `json:"created_at"`

	CommunityID   int    `json:"community_id"`
	CommunityName string `json:"community_name"`
	Background    string `json:"background"`
}

type TagInfoForm struct {
	TagName string `json:"tag_name"`
	TagType int    `json:"tag_type"`
}

type MediasForm struct {
	MediaURL  string `json:"media_url"`
	MediaType string `json:"media_type"`
}
