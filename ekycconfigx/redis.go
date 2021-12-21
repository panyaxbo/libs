package ekycconfigx

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/panyaxbo/libs/logx"
	"github.com/panyaxbo/libs/redisx"
)

type Redis struct {
	client *redisx.Client
}

func NewRedis(c *redisx.Client) *Redis {
	return &Redis{
		client: c,
	}
}

func (r *Redis) HGet(ctx context.Context, key, field string) (string, error) {
	b := r.client.HGet(ctx, key, field)
	return b.Val, b.Err
}

func (r *Redis) HSet(ctx context.Context, key string, data map[string]interface{}) error {
	return r.client.HSet(ctx, key, data)
}

func NewRedisClient(host string) *redisx.Client {
	logx.Infof("[CONIFG] [REDIS] addr:%s", host)

	return redisx.NewClient(&redis.Options{
		Addr: host,
		DB:   0,
	})
}
