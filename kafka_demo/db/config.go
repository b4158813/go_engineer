package db

import "fmt"

type SystemConfig struct {
	Kafka KafkaConfig `json:"kafka_config"`
}

type KafkaConfig struct {
	Url    string `json:"url"`
	Topic1 string `json:"topic_1"`
}

var C SystemConfig

func InitConfig() {
	kafka_config := KafkaConfig{
		Url:    "172.19.0.2:9092",
		Topic1: "mytest",
	}
	C.Kafka = kafka_config
	fmt.Printf("Init config done: %+v\n", C)
}
