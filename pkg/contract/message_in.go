package contract

import gosqs "github.com/engelmi/go-sqs"

type IncomingMessage struct {
	MessageId              *string                                   `json:"messageId"`
	ReceiptHandle          *string                                   `json:"receiptHandle"`
	Body                   *string                                   `json:"body"`
	MD5OfBody              *string                                   `json:"md5OfBody"`
	MessageAttributes      map[string]*IncomingMessageAttributeValue `json:"messageAttributes"`
	MD5OfMessageAttributes *string                                   `json:"md5OfMessageAttributes"`
	Attributes             map[string]*string                        `json:"attributes"`
}

type IncomingMessageAttributeValue struct {
	DataType    *string `json:"dataType"`
	BinaryValue []byte  `json:"binaryValue"`
	StringValue *string `json:"stringValue"`
}

func FromGoSqsMessage(msg gosqs.IncomingMessage) IncomingMessage {
	msgAttributes := map[string]*IncomingMessageAttributeValue{}
	for attribute, value := range msg.MessageAttributes {
		if value != nil {
			msgAttributes[attribute] = &IncomingMessageAttributeValue{
				DataType:    value.DataType,
				BinaryValue: value.BinaryValue,
				StringValue: value.StringValue,
			}
		}
	}

	return IncomingMessage{
		MessageId:              msg.MessageId,
		ReceiptHandle:          msg.ReceiptHandle,
		Body:                   msg.Body,
		MD5OfBody:              msg.MD5OfBody,
		MessageAttributes:      msgAttributes,
		MD5OfMessageAttributes: msg.MD5OfMessageAttributes,
		Attributes:             msg.Attributes,
	}
}
