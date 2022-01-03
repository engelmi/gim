package config

import (
	gosqs "github.com/engelmi/go-sqs"
	"github.com/sirupsen/logrus"
)

type GimConsumer struct {
	Consumer []Consumer `json:"consumer"`
}

type Consumer struct {
	Queue             `json:"queue"`
	ForwardUrl        string   `json:"forwardUrl"`
	ProcessingTimeout Duration `json:"processingTimeout"`
	PollTimeout       Duration `json:"pollTimeout"`
	AckTimeout        Duration `json:"ackTimeout"`
	BulkReadSize      int64    `json:"bulkReadSize"`
}

func (c Consumer) ToGoSqsConfig() gosqs.ConsumerConfig {
	return gosqs.ConsumerConfig{
		QueueConfig: gosqs.QueueConfig{
			Region:   c.Region,
			Endpoint: c.Endpoint,
			Queue:    c.QueueName,
		},
		PollTimeout:         c.PollTimeout.ToTimeDuration(),
		AckTimeout:          c.AckTimeout.ToTimeDuration(),
		MaxNumberOfMessages: c.BulkReadSize,
		Logger:              *logrus.New(),
	}
}
