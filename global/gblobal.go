package global

import (
	"finders-server/config"

	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	VP     *viper.Viper
	CONFIG *config.Server
	LOG    *oplogging.Logger
)
