package internal

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func InitRedis() {
	h := ViperConf.RedisConfig.Host
	p := ViperConf.RedisConfig.Port
	addr := fmt.Sprintf("%s:%d", h, p)
	fmt.Println(addr, "redis addr")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	var ctx = context.Background()

	err := rdb.Set(ctx, "key", "value", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis connected", val)
}
