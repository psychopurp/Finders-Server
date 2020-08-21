package responseForm

type GetMomentsResponseForm struct {
	Paginator
	NickName string             `json:"nick_name"`
	UserID   string             `json:"user_id"`
	Avatar   string             `json:"avatar"`
	Moments  []SimpleMomentForm `json:"moments"`
}

type SimpleMomentForm struct {
	MomentID   string       `json:"moment_id"`
	MomentInfo string       `json:"moment_info"`
	ReadNum    int          `json:"read_num"`
	Medias     []MediasForm `json:"medias"`
	CreatedAt  string       `json:"created_at"`
}
