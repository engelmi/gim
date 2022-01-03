package contract

import gosqs "github.com/engelmi/go-sqs"

type OutgoingMessage struct {
	DeduplicationId *string           `json:"deduplicationId"`
	GroupId         *string           `json:"groupId"`
	Payload         string            `json:"payload"`
	Attributes      map[string]string `json:"attributes"`
}

func (m OutgoingMessage) ToGoSqsMessage() gosqs.OutgoingMessage {
	return gosqs.OutgoingMessage{
		DeduplicationId: m.DeduplicationId,
		GroupId:         m.GroupId,
		Payload:         []byte(m.Payload),
		Attributes:      m.Attributes,
	}
}
