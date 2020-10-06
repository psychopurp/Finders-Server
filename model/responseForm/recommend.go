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
