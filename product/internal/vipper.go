package internal

import (
	"encoding/json"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

var ViperConf ViperConfig
var ConfigFilePath = "../dev_config.yaml"
var NacosConf NacosContent

type NacosContent struct {
	NacosConfig NacosConfig `json:"nacos" mapstructure:"nacos"`
}

func init() {
	initNacos()
	initFromNacos()
	InitDB()
	InitRedis()
	fmt.Println("Config Initial Successfully!")
}

func initNacos() {
	v := viper.New()
	configName := ConfigFilePath
	v.SetConfigFile(configName)
	v.ReadInConfig()
	v.Unmarshal(&NacosConf)
}

func initFromNacos() {
	nacosConfig := NacosConf.NacosConfig

	serverConfigs := []constant.ServerConfig{
		{IpAddr: nacosConfig.Host, Port: uint64(nacosConfig.Port)},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConfig.Namespace, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "nacos/log",
		CacheDir:            "nacos/cache",
		LogLevel:            "debug",
	}

	// Create config client for dynamic configuration
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group})

	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(content), &ViperConf)
	fmt.Println(ViperConf, "content")
}

type ViperConfig struct {
	RedisConfig    RedisConfig    `mapstructure:"redis" json:"redis"`
	ConsulConfig   ConsulConfig   `mapstructure:"consul" json:"consul"`
	DBConfig       DBConfig       `mapstructure:"mysql" json:"mysql"`
	ProductSrvConf ProductSrvConf `mapstructure:"product_srv" json:"product_srv"`
	ProductWebConf ProductWebConf `mapstructure:"product_web" json:"product_web"`
	JWTConfig      JWTConfig      `mapstructure:"jwt" json:"jwt"`
}
