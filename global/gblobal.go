package global

import (
	"finders-server/config"
<<<<<<< HEAD
=======
	"github.com/gomodule/redigo/redis"
>>>>>>> test

	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	VP     *viper.Viper
	CONFIG *config.Server
	LOG    *oplogging.Logger
	DB     *gorm.DB
<<<<<<< HEAD
=======
	RedisConn *redis.Pool
>>>>>>> test
)
