package pubsubx

import (
	"context"
	"encoding/json"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/panyaxbo/libs/configx"
	"github.com/panyaxbo/libs/logx"
	"github.com/panyaxbo/libs/maskx"
	"github.com/panyaxbo/libs/utilx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type PubConfig struct {
	Client    *pubsub.Client
	Timeout   time.Duration
	ProjectID string
}

type Pub struct {
	Client    *pubsub.Client
	Timeout   time.Duration
	ProjectID string
	Config    *configx.ConfigMaskLog
}

func NewPub(cfg PubConfig) *Pub {
	if cfg.Client == nil || cfg.Timeout == 0 || cfg.ProjectID == "" {
		logx.Panic("pubsub client, timeout, topicID, projectID required")
	}

	return &Pub{
		Client:    cfg.Client,
		Timeout:   cfg.Timeout,
		ProjectID: cfg.ProjectID,
	}
}

func NewPubWithMaskLog(cfg PubConfig, config *configx.ConfigMaskLog) *Pub {
	if cfg.Client == nil || cfg.Timeout == 0 || cfg.ProjectID == "" {
		logx.Panic("pubsub client, timeout, topicID, projectID required")
	}

	return &Pub{
		Client:    cfg.Client,
		Timeout:   cfg.Timeout,
		ProjectID: cfg.ProjectID,
		Config:    config,
	}
}

func (p *Pub) Publish(ctx context.Context, topicID string, i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return errors.WithStack(err)
	}

	pubCtx, cancel := context.WithTimeout(context.Background(), p.Timeout)
	defer cancel()

	topic := p.Client.Topic(topicID)

	start := time.Now()
	serverID, err := topic.Publish(pubCtx, &pubsub.Message{Data: b}).Get(pubCtx)

	logx.WithContext(ctx).WithFields(logrus.Fields{
		"topicID":   topicID,
		"projectID": p.ProjectID,
		"value":     logx.LimitMSGByte(b),
		"duration":  time.Since(start).String(),
		"serverID":  serverID,
	}).Info("pub information")

	return errors.WithStack(err)
}

func (p *Pub) logValue(b []byte) string {
	if p.Config == nil {
		return logx.LimitMSGByte(b)
	}

	m := maskx.Init(configx.SensitiveFields)

	if p.Config.IsMaskLogWithEncrypt {
		s, err := m.JsonMaskEncrypted(b, p.Config.Env)
		if err != nil {
			panic(err)
		}
		return logx.LimitMSGByte(utilx.StringPointerToByteArray(s))
	}

	if p.Config.IsMaskLogWithSymbol {
		s, err := m.JsonMaskSymbol(b, p.Config.Symbol)
		if err != nil {
			panic(err)
		}
		return logx.LimitMSGByte(utilx.StringPointerToByteArray(s))
	}

	return logx.LimitMSGByte(b)
}
