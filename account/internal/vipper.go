package internal

import (
	"github.com/spf13/viper"
	"ocean_mall/account/conf"
)

var ViperConf ViperConfig
var ConfigFilePath = "./dev_config.yaml"

func init() {
	v := viper.New()
	configName := ConfigFilePath
	v.SetConfigFile(configName)
	v.ReadInConfig()
	v.Unmarshal(&ViperConf)
}

type ViperConfig struct {
	RedisConfig    RedisConfig         `mapstructure:"redis"`
	ConsulConfig   ConsulConfig        `mapstructure:"consul"`
	DBConfig       DBConfig            `mapstructure:"mysql"`
	AccountSrvConf conf.AccountSrvConf `mapstructure:"account_srv"`
	AccountWebConf conf.AccountWebConf `mapstructure:"account_web"`
}
