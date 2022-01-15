package producer

import (
	"net/http"

	"github.com/engelmi/gim/internal/logger"
	"github.com/engelmi/gim/pkg/config"
	"github.com/engelmi/gim/pkg/contract"
	gosqs "github.com/engelmi/go-sqs"
)

func Handler(producer gosqs.Producer, config config.Producer) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		l := logger.GetLogger().WithFields(map[string]interface{}{
			"queue": map[string]interface{}{
				"name":     config.QueueName,
				"endpoint": config.Endpoint,
			},
			"producer-name": config.ProducerName,
		})
		l.Debug("Received request")

		var msg contract.OutgoingMessage
		err := decodeJsonBody(rw, r, &msg)
		if err != nil {
			l.WithError(err).Error("Failed to decode json body of request")

			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Could not decode request body"))
			return
		}

		msgIdPtr, err := producer.Send(r.Context(), msg.ToGoSqsMessage())
		if err != nil {
			l.WithError(err).Error("Failed to send message")

			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}

		msgId := ""
		if msgIdPtr != nil {
			msgId = *msgIdPtr
		}

		l.WithField("message-id", msgId).Debug("Send message successfully")

		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte(msgId))
	}
}
