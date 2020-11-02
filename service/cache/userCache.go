package cache

import (
	"encoding/json"
	"finders-server/model"
	gredis2 "finders-server/pkg/gredis"
	"finders-server/st"
	"fmt"
	"strings"
)

type UserCacheService struct {}

func NewUserCacheService()*UserCacheService{
	return &UserCacheService{}
}

func (s *UserCacheService)GetUserCacheKey() string{
	keys := []string{
		CACHE_TAG,
		CACHE_USER,
	}
	return  strings.Join(keys, SEP)
}

func (s *UserCacheService)GetUserByUserId(userId string)(user model.User, err error){
	var(
		data []byte
	)
	key := getCacheKey(s.GetUserCacheKey(), userId)
	if gredis2.Exists(key){
		data, err = gredis2.Get(key)
		if err != nil {
			st.Debug("user cache error", err)
			return
		} else{
			_ = json.Unmarshal(data, &user)
			return user, nil
		}
	}
	return user, fmt.Errorf("cache not find")
}

func (s *UserCacheService)SetUserByUserId(user model.User)(err error){
	keyBase := s.GetUserCacheKey()
	keys := []string{
		keyBase,
		user.UserID.String(),
	}
	key := strings.Join(keys, SEP)
	err = gredis2.Set(key, user, 3600)
	return
}