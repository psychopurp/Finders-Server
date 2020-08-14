package requestForm

type CreateActivityForm struct {
	CommunityID  int      `json:"community_id" validate:"required,gte=0"`
	ActivityInfo string   `json:"activity_info" validate:"required,min=1,max=65535"`
	MediaIDs     []string `json:"media_ids" validate:"required,min=1,max=50"`
}

type GetActivityIDForm struct {
	ActivityID string `json:"activity_id" validate:"required,min=1,max=50"`
}
