package consumer

import (
	"context"

	"github.com/engelmi/gim/internal/logger"
	"github.com/engelmi/gim/pkg/config"
	"github.com/engelmi/gim/pkg/contract"
	gosqs "github.com/engelmi/go-sqs"
)

func Handler(client ApiClient, config config.Consumer) gosqs.MessageHandler {
	return func(ctx context.Context, msg gosqs.IncomingMessage) error {
		l := logger.GetLogger().WithFields(map[string]interface{}{
			"message-id": msg.MessageId,
			"queue": map[string]interface{}{
				"name":     config.QueueName,
				"endpoint": config.Endpoint,
			},
			"forward-url": config.ForwardUrl,
		})
		l.Debug("Received message")

		incomingMsg := contract.FromGoSqsMessage(msg)
		err := client.Send(incomingMsg)
		if err != nil {
			l.WithError(err).Error("Failed to send message")
			return err
		}
		return nil
	}
}
