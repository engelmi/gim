package producer

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/engelmi/gim/internal/logger"
	"github.com/engelmi/gim/pkg/config"
	gosqs "github.com/engelmi/go-sqs"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type Producer struct {
	server http.Server
}

func NewProducer(gimProducerConfig config.GimProducer) (Producer, error) {
	router := mux.NewRouter()
	for _, producerConfig := range gimProducerConfig.Producer {
		sqsProducer, err := gosqs.NewProducer(producerConfig.ToGoSqsConfig())
		if err != nil {
			return Producer{}, errors.Wrap(err, "Could not create sqs producer")
		}
		router.Methods(http.MethodPost).Path(fmt.Sprintf("/%s/produce", producerConfig.ProducerName)).HandlerFunc(Handler(sqsProducer, producerConfig))
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", gimProducerConfig.Server.Port),
		Handler: router,
	}

	logger.GetLogger().Info("Producer set up successfully")

	return Producer{
		server: server,
	}, nil
}

func (s Producer) Start(ctx context.Context, wg *sync.WaitGroup) {
	go s.server.ListenAndServe()
	logger.GetLogger().Info(fmt.Sprintf("Started producer on port %s", s.server.Addr))

	go func() {
		<-ctx.Done()

		ctxWithTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		logger.GetLogger().Info("Shutting down producer...")
		s.server.Shutdown(ctxWithTimeout)
		logger.GetLogger().Info("Finished shutting down producer")

		wg.Done()
	}()

	wg.Add(1)
}
