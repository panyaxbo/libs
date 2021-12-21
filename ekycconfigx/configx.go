package ekycconfigx

import (
	"context"
	"errors"
	"strings"

	"github.com/panyaxbo/libs/logx"
)

type EkycConfig struct {
	env string
}

func NewEkycConfig(env string) *EkycConfig {
	return &EkycConfig{env}
}

func (e *EkycConfig) GetEKYCConfigValueByKey(ctx context.Context, key string) (string, error) {
	if e.env == "" {
		return "", errors.New("Please specify environment to get config from redis")
	}
	host := e.checkHostFromEnvironment(e.env)

	logx.WithSeverityInfo(ctx).Infof("host %s", host)

	r := newRedis(newRedisClient(host))

	val, err := r.HGet(ctx, ekycDB, key)
	if err != nil {
		return "", err
	}

	return val, nil
}

func (e *EkycConfig) GetDipchipConfigValueByKey(ctx context.Context, key string) (string, error) {
	if e.env == "" {
		return "", errors.New("Please specify environment to get config from redis")
	}
	host := e.checkHostFromEnvironment(e.env)

	logx.WithSeverityInfo(ctx).Infof("host %s", host)

	r := newRedis(newRedisClient(host))

	val, err := r.HGet(ctx, dipchipDB, key)

	if err != nil {
		return "", err
	}

	return val, nil
}

func (e *EkycConfig) checkHostFromEnvironment(env string) string {
	host := ""
	if strings.EqualFold(e.env, "loc") || strings.EqualFold(e.env, "local") {
		host = redisHost["redisHostLocal"]
	} else if strings.EqualFold(e.env, "dev") || strings.EqualFold(e.env, "develop") {
		host = redisHost["redisHostDev"]
	} else if strings.EqualFold(e.env, "uat") {
		host = redisHost["redisHostUat"]
	} else if strings.EqualFold(e.env, "prd") || strings.EqualFold(e.env, "prod") || strings.EqualFold(e.env, "production") {
		host = redisHost["redisHostProd"]
	}

	return host
}
