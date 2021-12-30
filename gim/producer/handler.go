package producer

import (
	"net/http"

	gosqs "github.com/engelmi/go-sqs"
)

func Handler(producer gosqs.Producer) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var msg OutgoingMessage
		err := decodeJsonBody(rw, r, &msg)
		if err != nil {
			// todo: logging
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Could not decode request body"))
			return
		}

		msgId, err := producer.Send(r.Context(), msg.ToGoSqsMessage())
		if err != nil {
			// todo: logging
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte(*msgId))
	}
}
