package internal

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	DataId    string `mapstructure:"dataId" json:"dataId"`
	Group     string `mapstructure:"group" json:"group"`
}
