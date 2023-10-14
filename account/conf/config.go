package conf

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	JWTConfig JWTConfig `mapstructure:"jwt"`
}

var AppConf AppConfig

func init() {
	v := viper.New()
	configName := "dev_config.yaml"
	v.SetConfigFile(configName)
	v.ReadInConfig()
	v.Unmarshal(&AppConf)
}
