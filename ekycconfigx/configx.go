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

func (e *EkycConfig) GetConfigValueByKey(ctx context.Context, env, key string) error {
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

	logx.WithSeverityInfo(ctx).Infox("host %s", host)

	return nil
}
