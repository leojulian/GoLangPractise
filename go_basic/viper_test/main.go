package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type MySqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	MySql       MySqlConfig `mapstructure:"mysql"`
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main() {
	v := viper.New()
	if GetEnvInfo("leoshopdebug") {
		v.SetConfigFile("config-debug.yaml")
	} else {
		v.SetConfigFile("config-pro.yaml")
	}
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}

	fmt.Println(serverConfig)

	// watch config change
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serverConfig)
		fmt.Println(serverConfig)
	})

	time.Sleep(time.Second * 300)
}
