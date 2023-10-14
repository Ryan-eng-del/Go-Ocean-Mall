package conf

type JWTConfig struct {
	SignKey string `mapstructure:"sign_key"`
}
