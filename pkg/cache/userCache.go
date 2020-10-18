package cache

import (
	"encoding/json"
	"finders-server/model"
	"finders-server/service/gredis"
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
	keyBase := s.GetUserCacheKey()
	keys := []string{
		keyBase,
		userId,
	}
	key := strings.Join(keys, SEP)
	if gredis.Exists(key){
		data, err = gredis.Get(key)
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
	err = gredis.Set(key, user, 3600)
	return
}