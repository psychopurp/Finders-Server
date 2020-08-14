package responseForm

import uuid "github.com/satori/go.uuid"

type SimpleUserInfo struct {
	UserId       uuid.UUID
	Avatar       string
	NickName     string
	Introduction string
}
