package internal

import (
	"context"
	"sync"

	"github.com/engelmi/gim/internal/consumer"
	"github.com/engelmi/gim/internal/logger"
	"github.com/engelmi/gim/internal/producer"
	"github.com/engelmi/gim/pkg/config"
)

type GopherInTheMiddle struct {
	gimProducer producer.Producer
	gimConsumer consumer.Consumer
}

func NewGopherInTheMiddle(gimConfig config.Gim) (GopherInTheMiddle, error) {
	producer, err := producer.NewProducer(gimConfig.GimProducer)
	if err != nil {
		return GopherInTheMiddle{}, err
	}

	consumer, err := consumer.NewConsumer(gimConfig.GimConsumer)
	if err != nil {
		return GopherInTheMiddle{}, err
	}

	return GopherInTheMiddle{
		gimProducer: producer,
		gimConsumer: consumer,
	}, nil
}

func (g GopherInTheMiddle) Start(ctx context.Context, wg *sync.WaitGroup) {
	l := logger.GetLogger()

	if wg == nil {
		wg = &sync.WaitGroup{}
	}

	g.gimProducer.Start(ctx, wg)
	g.gimConsumer.Start(ctx, wg)

	l.Info("Gophers are ready and waiting to work")
	wg.Wait()
}
