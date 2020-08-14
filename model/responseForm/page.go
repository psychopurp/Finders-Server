package responseForm

type Paginator struct {
	TotalCNT  int `json:"total_cnt"`
	TotalPage int `json:"total_page"`
	CNT       int `json:"cnt"`
	Page      int `json:"page"`
}
