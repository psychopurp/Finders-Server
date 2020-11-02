package cache

import (
	"encoding/json"
	"finders-server/pkg/gredis"
	"finders-server/pkg/smsApi"
	"finders-server/st"
	"finders-server/utils/reg"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type PhoneCacheService struct {
	Phone string
}

func NewPhoneCacheService() *PhoneCacheService {
	return &PhoneCacheService{}
}

// GetPhoneCodeKey 获取手机号验证码的 key
func (s *PhoneCacheService) GetPhoneCodeKey() string {
	keys := []string{
		CACHE_TAG,
		CACHE_PHONE,
	}
	return strings.Join(keys, SEP)
}

func (s *PhoneCacheService) GenerateCode() string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < 6; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func (s *PhoneCacheService) GetPhoneCode(phone string) (code string, err error) {
	if !reg.Phone(phone) {
		return "", fmt.Errorf("非法手机号")
	}
	key := getCacheKey(s.GetPhoneCodeKey(), phone)
	keyBool := getCacheKey(s.GetPhoneCodeKey(), phone+"bool")
	if gredis.Exists(keyBool) {
		return "", fmt.Errorf("请勿重复发送短信验证码，1分钟后再试")
	}
	code = s.GenerateCode()
	st.Debug("code:", code)
	err = smsApi.SendSMS(phone, code)
	if err != nil {
		return
	}
	err = gredis.Set(key, code, 360)
	if err != nil {
		st.Debug("set redis error ", err)
		return "", fmt.Errorf("发送短信验证码失败")
	}
	err = gredis.Set(keyBool, code, 60)
	if err != nil {
		st.Debug("set redis error2 ", err)
		return "", fmt.Errorf("发送短信验证码失败")
	}
	return
}

func (s *PhoneCacheService) ValidatePhoneCode(phone, code string) bool {
	var (
		err      error
		data     []byte
		realCode string
	)
	key := getCacheKey(s.GetPhoneCodeKey(), phone)
	if !gredis.Exists(key) {
		return false
	}
	data, err = gredis.Get(key)

	if err != nil {
		st.Debug("get redis error")
		return false
	}
	err = json.Unmarshal(data, &realCode)
	if err != nil {
		st.Debug("unmarshal error", err)
		return false
	}
	//st.Debug("real code:", realCode)
	if code == realCode{
		_, err = gredis.Delete(key)
		if err != nil {
			st.Debug("delete code error", err)
		}
		return true
	}
	return false
}
