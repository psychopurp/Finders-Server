package initialize

import (
	"finders-server/global"
	"finders-server/st"
	"github.com/gomodule/redigo/redis"
	"time"
)

func Redis() error {
	st.Debug("address", global.CONFIG.Redis.Addr)
	global.RedisConn = &redis.Pool{
		MaxIdle:     global.CONFIG.Redis.MaxIdle,
		MaxActive:   global.CONFIG.Redis.MaxActive,
		IdleTimeout: global.CONFIG.Redis.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", global.CONFIG.Redis.Addr)
			if err != nil {
				return nil, err
			}
			if global.CONFIG.Redis.Password != "" {
				if _, err := c.Do("AUTH", global.CONFIG.Redis.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}