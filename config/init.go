package config

import (
	"fmt"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 加载配置文件
func InitConfig(configDir string) (config *Server) {
	v := viper.New()
	configFile, err := filepath.Glob(configDir + "/config*.yaml")
	if err != nil || len(configFile) <= 0 {
		panic("config file not found")
	}
	v.SetConfigFile(configFile[0])
	err = v.ReadInConfig()
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
