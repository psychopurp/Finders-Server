package responseForm

type SimpleUserInfo struct {
	UserId       string `json:"user_id"`
	Avatar       string `json:"avatar"`
	NickName     string `json:"nickname"`
	Introduction string `json:"introduction"`
	Signature    string `json:"signature"`
}

type SimpleUserInfoWithPage struct {
	Paginator
	SimpleUserInfos []SimpleUserInfo `json:"simple_user_infos"`
}
