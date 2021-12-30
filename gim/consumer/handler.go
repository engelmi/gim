package consumer

import (
	"context"

	gosqs "github.com/engelmi/go-sqs"
)

func Handler(client ApiClient) gosqs.MessageHandler {
	return func(ctx context.Context, msg gosqs.IncomingMessage) error {
		incomingMsg := FromGoSqsMessage(msg)
		err := client.Send(incomingMsg)
		if err != nil {
			return err
		}
		return nil
	}
}
