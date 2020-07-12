package cache_service

type Phone struct {
	Phone string
}

const (
	CACHE_PHONE = "PHONE"
)
// GetPhoneCodeKey 获取手机号验证码的 key
func (p Phone) GetPhoneCodeKey() string {
	return CACHE_PHONE + "_" + p.Phone
}
