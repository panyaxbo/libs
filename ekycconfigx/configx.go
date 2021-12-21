package ekycconfigx

import (
	"context"

	"github.com/panyaxbo/libs/logx"
)

type EkycConfig struct {
	// redisHost string
	// redisPort string
	env string
	key string
}

func NewEkycConfig(env string, key string) *EkycConfig {
	return &EkycConfig{env, key}
}

func (e *EkycConfig) GetEKYCConfigValueByKey(ctx context.Context, env, key string) (string, error) {
	host := ""
	if env == "loc" {
		host = redisHost["redisHostLocal"]
	} else if env == "dev" {
		host = redisHost["redisHostDev"]
	} else if env == "uat" {
		host = redisHost["redisHostUat"]
	} else if env == "prd" {
		host = redisHost["redisHostProd"]
	}

	logx.WithSeverityInfo(ctx).Infof("host %s", host)

	r := NewRedis(NewRedisClient(host))

	val, err := r.HGet(ctx, ekycDB, key)
	if err != nil {
		return "", err
	}

	return val, nil
}

func (e *EkycConfig) GetDipchipConfigValueByKey(ctx context.Context, env, key string) (string, error) {
	host := ""
	if env == "loc" {
		host = redisHost["redisHostLocal"]
	} else if env == "dev" {
		host = redisHost["redisHostDev"]
	} else if env == "uat" {
		host = redisHost["redisHostUat"]
	} else if env == "prd" {
		host = redisHost["redisHostProd"]
	}

	logx.WithSeverityInfo(ctx).Infof("host %s", host)

	r := NewRedis(NewRedisClient(host))

	val, err := r.HGet(ctx, dipchipDB, key)

	if err != nil {
		return "", err
	}

	return val, nil
}
