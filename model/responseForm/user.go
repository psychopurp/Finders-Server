package responseForm

type SimpleUserInfo struct {
	UserId       string `json:"user_id"`
	Avatar       string    `json:"avatar"`
	NickName     string    `json:"nick_name"`
	Introduction string    `json:"introduction"`
	Signature    string    `json:"signature"`
}
