package requestForm

type CreateCommunityForm struct {
	CommunityName        string `json:"community_name" validate:"required,min=1,max=100"`
	CommunityDescription string `json:"community_description" validate:"required,min=1,max=65535"`
	Background           string `json:"background" validate:"required,min=1,max=200"`
}

type UpdateCommunityForm struct {
	CommunityID          int    `json:"community_id" validate:"required,gte=0"`
	CommunityName        string `json:"community_name" validate:"omitempty,min=1,max=100"`
	CommunityDescription string `json:"community_description" validate:"omitempty,min=1,max=65535"`
	Background           string `json:"background" validate:"omitempty,min=1,max=200"`
}

type GetCommunityIDForm struct {
	CommunityID int `json:"community_id" validate:"required,gte=0"`
}
