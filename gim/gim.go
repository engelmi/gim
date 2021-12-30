package gim

import (
	"context"
	"sync"

	"github.com/engelmi/gim/config"
	"github.com/engelmi/gim/gim/consumer"
	"github.com/engelmi/gim/gim/producer"
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
	if wg == nil {
		wg = &sync.WaitGroup{}
	}

	g.gimProducer.Start(ctx, wg)
	g.gimConsumer.Start(ctx, wg)

	wg.Wait()
}
