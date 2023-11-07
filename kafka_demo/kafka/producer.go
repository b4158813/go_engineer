package kafka

import (
	"context"
	"encoding/json"
	"my_test/cn"
	"my_test/db"
	"my_test/utils"

	"github.com/IBM/sarama"
)

var kafkaConfig *sarama.Config = sarama.NewConfig()

func ProduceMessages(ctx context.Context, topicName string, structMsg interface{}) error {
	ret, err := json.Marshal(structMsg)
	if err != nil {
		db.L.Errorf("[%s] json.Marshal error: %s\n", utils.GetCurrentFunctionName(), err)
		return err
	}

	client, err := sarama.NewSyncProducer([]string{db.C.Kafka.Url}, kafkaConfig)
	if err != nil {
		db.L.Errorf("[%s] sarama.NewSyncProducer error: %s\n", utils.GetCurrentFunctionName(), err)
		return err
	}
	defer client.Close()

	msg := &sarama.ProducerMessage{
		Topic: db.C.Kafka.Topic1,
		Key:   sarama.StringEncoder(cn.Schema + cn.Name),
		Value: sarama.ByteEncoder(ret),
	}
	_, _, err = client.SendMessage(msg)
	if err != nil {
		db.L.Errorf("[%s] client.SendMessage error: %s\n", utils.GetCurrentFunctionName(), err)
		return err
	}
	return nil
}

func InitKafkaCfg() {
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	kafkaConfig.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
}
