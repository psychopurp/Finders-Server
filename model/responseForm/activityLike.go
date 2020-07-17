package responseForm

type ActivityLikesResponseForm struct {
	Paginator
	LikeForms []LikeForm `json:"likes"`
}

type LikeForm struct {
	NickName string `json:"nick_name"`
	UserID   string `json:"user_id"`
	Avatar   string `json:"avatar"`
}
