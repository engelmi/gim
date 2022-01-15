package consumer

import (
	"context"
	"sync"

	"github.com/engelmi/gim/internal/logger"
	"github.com/engelmi/gim/pkg/config"
	gosqs "github.com/engelmi/go-sqs"
	"github.com/pkg/errors"
)

type Consumer struct {
	consumer []gosqs.Consumer
}

func NewConsumer(gimConsumerConfig config.GimConsumer) (Consumer, error) {
	consumer := make([]gosqs.Consumer, 0, len(gimConsumerConfig.Consumer))
	for _, gimConsumerConfig := range gimConsumerConfig.Consumer {
		apiClient := NewApiClient(gimConsumerConfig.ForwardUrl, gimConsumerConfig.ProcessingTimeout.ToTimeDuration())
		sqsConsumer, err := gosqs.NewConsumer(gimConsumerConfig.ToGoSqsConfig(), Handler(apiClient, gimConsumerConfig))
		if err != nil {
			return Consumer{}, errors.Wrap(err, "Could not create sqs consumer")
		}
		consumer = append(consumer, sqsConsumer)
	}

	logger.GetLogger().Info("Consumer set up successfully")

	return Consumer{
		consumer: consumer,
	}, nil
}

func (l Consumer) Start(ctx context.Context, wg *sync.WaitGroup) {
	for _, consumer := range l.consumer {
		go consumer.StartListening(ctx, wg)
	}
	logger.GetLogger().Info("Started all consumer")
}
