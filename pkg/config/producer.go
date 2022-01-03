package config

import gosqs "github.com/engelmi/go-sqs"

type GimProducer struct {
	Server   HttpServer `json:"server"`
	Producer []Producer `json:"producer"`
}

type Producer struct {
	Queue        `json:"queue"`
	ProducerName string   `json:"producerName"`
	SendTimeout  Duration `json:"sendTimeout"`
}

func (p Producer) ToGoSqsConfig() gosqs.ProducerConfig {
	return gosqs.ProducerConfig{
		QueueConfig: gosqs.QueueConfig{
			Region:   p.Region,
			Endpoint: p.Endpoint,
			Queue:    p.QueueName,
		},
		Timeout: p.SendTimeout.ToTimeDuration(),
	}
}
