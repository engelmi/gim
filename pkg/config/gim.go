package config

import (
	"encoding/json"
	"os"
)

type Gim struct {
	GimConsumer GimConsumer `json:"gimconsumer"`
	GimProducer GimProducer `json:"gimproducer"`
}

func (gim Gim) ToFile(name string) error {
	data, err := json.Marshal(gim)
	if err != nil {
		return err
	}
	return os.WriteFile(name, data, os.ModePerm)
}

func FromFile(name string) (Gim, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return Gim{}, err
	}
	var gimConfig Gim
	err = json.Unmarshal(data, &gimConfig)
	if err != nil {
		return Gim{}, err
	}
	return gimConfig, nil
}

func FromJsonString(gimStr string) (Gim, error) {
	var gimConfig Gim
	err := json.Unmarshal([]byte(gimStr), &gimConfig)
	if err != nil {
		return Gim{}, err
	}
	return gimConfig, nil
}

type Queue struct {
	Region    string `json:"region"`
	Endpoint  string `json:"endpoint"`
	QueueName string `json:"queueName"`
}

type HttpServer struct {
	Port string `json:"port"`
}
