package responseForm

type CommentResponseForm struct {
	Paginator
	CommentForms []CommentInfoForm `json:"comments"`
}

type CommentInfoForm struct {
	CommentID int    `json:"comment_id"`
	Content   string `json:"content"`
	NickName  string `json:"nick_name"`
	UserID    string `json:"user_id"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
	ReplyNum  int    `json:"reply_num"`
}
