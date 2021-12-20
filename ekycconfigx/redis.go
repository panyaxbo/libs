package ekycconfigx

import (
	"github.com/go-redis/redis"
	"github.com/panyaxbo/libs/logx"
	"github.com/panyaxbo/libs/redisx"
	"github.com/spf13/viper"
)

type Redis struct {
	client *redisx.Client
}

func NewRedis(c *redisx.Client) *Redis {
	return &Redis{
		client: c,
	}
}

func NewRedisClient() *redisx.Client {
	logx.Infof("[CONIFG] [REDIS] addr:%s", viper.GetString("redis.addr"))

	return redisx.NewClient(&redis.Options{
		Addr: viper.GetString("redis.addr"),
		DB:   0,
	})
}
