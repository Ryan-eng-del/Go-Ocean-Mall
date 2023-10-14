package conf

import (
	"github.com/spf13/viper"
)

type AccountSrvConf struct {
	Host    string   `mapstructure:"host"`
	Port    string   `mapstructure:"port"`
	SrvName string   `mapstructure:"srvName"`
	Tags    []string `mapstructure:"tags"`
}

type AccountWebConf struct {
	Host    string   `mapstructure:"host"`
	Port    string   `mapstructure:"port"`
	SrvName string   `mapstructure:"srvName"`
	Tags    []string `mapstructure:"tags"`
}

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
