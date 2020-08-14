package responseForm

type UploadResponseForm struct {
	Medias []UploadMedia `json:"medias"`
}

type UploadMedia struct {
	MediaID  string `json:"media_id"`
	MediaURL string `json:"media_url"`
}
