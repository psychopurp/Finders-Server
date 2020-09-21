package responseForm

type GetUserMomentsResponseForm struct {
	Paginator
	NickName  string   `json:"nick_name"`
	UserID    string   `json:"user_id"`
	Avatar    string   `json:"avatar"`
	MomentIDs []string `json:"moment_ids"`
}

type GetMomentResponseForm struct {
	NickName   string       `json:"nick_name"`
	Avatar     string       `json:"avatar"`
	UserID     string       `json:"user_id"`
	Signature  string       `json:"signature"`
	MomentID   string       `json:"moment_id"`
	MomentInfo string       `json:"moment_info"`
	Location   string       `json:"location"`
	ReadNum    int          `json:"read_num"`
	LikeNum    int          `json:"like_num"`
	Medias     []MediasForm `json:"medias"`
	CreatedAt  string       `json:"created_at"`
}
