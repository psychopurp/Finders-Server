package core

import (
	"finders-server/global"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

// 加载配置文件
func init() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	//如果文件改变
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed...")
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Printf("Unmarshal err: %s\n", err)
		}
	})

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Printf("Unmarshal err: %s\n", err)
	}
	global.VP = v
}
