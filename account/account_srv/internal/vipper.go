package internal

import (
	"github.com/spf13/viper"
)

var ViperConf ViperConfig

func init() {
	v := viper.New()
	configName := "dev_config.yaml"
	v.SetConfigFile(configName)
	v.ReadInConfig()
	v.Unmarshal(&ViperConf)
}

type ViperConfig struct {
	RedisConfig RedisConfig `mapstructure:"redis"`
}
