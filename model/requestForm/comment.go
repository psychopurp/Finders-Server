package requestForm

type CreateCommentForm struct {
	ItemID  string `json:"item_id" validate:"required,min=1,max=50"`
	Content string `json:"content" validate:"required,min=1,max=65535"`
}
