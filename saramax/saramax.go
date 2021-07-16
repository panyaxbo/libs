package saramax

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/panyaxbo/libs/logx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Config struct {
	*sarama.Config
	Brokers []string
}

type Produce struct {
	SyncProducer sarama.SyncProducer
}

func NewProduce(cfg *Config) *Produce {
	syncProducer, err := sarama.NewSyncProducer(cfg.Brokers, cfg.Config)
	if err != nil {
		logx.Panic(err)
	}

	return &Produce{
		SyncProducer: syncProducer,
	}
}

func (p *Produce) Produce(ctx context.Context, topic string, i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return errors.WithStack(err)
	}

	return p.produce(ctx, topic, b)
}

func (p *Produce) ProduceWithHeaders(ctx context.Context, topic string, i interface{}, h map[string]string) error {
	b, err := json.Marshal(i)
	if err != nil {
		return errors.WithStack(err)
	}

	return p.produceWithHeaders(ctx, topic, b, h)
}

func (p *Produce) ProduceNoMarshal(ctx context.Context, topic string, b []byte) error {
	return p.produce(ctx, topic, b)
}

func (p *Produce) produce(ctx context.Context, topic string, b []byte) error {
	start := time.Now()
	partition, offset, err := p.SyncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(b),
	})

	logx.WithContext(ctx).WithFields(logrus.Fields{
		"value":     logx.LimitMSGByte(b),
		"topic":     topic,
		"partition": partition,
		"offset":    offset,
		"duration":  time.Since(start).String(),
	}).Info("produce information")

	return errors.WithStack(err)
}

func (p *Produce) produceWithHeaders(ctx context.Context, topic string, b []byte, h map[string]string) error {
	start := time.Now()
	var sh []sarama.RecordHeader
	for key, value := range h {
		sh = append(sh, sarama.RecordHeader{
			Key:   []byte(key),
			Value: []byte(value),
		})
	}

	partition, offset, err := p.SyncProducer.SendMessage(&sarama.ProducerMessage{
		Topic:   topic,
		Value:   sarama.ByteEncoder(b),
		Headers: sh,
	})

	logx.WithContext(ctx).WithFields(logrus.Fields{
		"headers":   fmt.Printf("%+v", h),
		"value":     logx.LimitMSGByte(b),
		"topic":     topic,
		"partition": partition,
		"offset":    offset,
		"duration":  time.Since(start).String(),
	}).Info("produce information")

	return errors.WithStack(err)
}
