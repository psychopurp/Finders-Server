package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

// 加载配置文件
func InitConfig() (config *Server) {
	v := viper.New()

	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	v.WatchConfig()

	//如果文件改变
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed...")
		if err := v.Unmarshal(&config); err != nil {
			fmt.Printf("Unmarshal err: %s\n", err)
		}
	})

	if err := v.Unmarshal(&config); err != nil {
		fmt.Printf("Unmarshal err: %s\n", err)
	}
	return
}
