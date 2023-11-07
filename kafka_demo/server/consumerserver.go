package server

import (
	"fmt"
	"my_test/db"
	"my_test/utils"

	"github.com/IBM/sarama"
)

func StartConsumerServer() {
	consumer, err := sarama.NewConsumer([]string{db.C.Kafka.Url}, nil)
	if err != nil {
		db.L.Fatalf("[%s] sarama.NewConsumer error: %s\n", utils.GetCurrentFunctionName(), err)
	}
	defer consumer.Close()

	partitions, err := consumer.Partitions(db.C.Kafka.Topic1)
	if err != nil {
		db.L.Fatalf("[%s] consumer.Partitions error: %s\n", utils.GetCurrentFunctionName(), err)
	}
	fmt.Println("start consumer server done!")

	for {
		for _, p := range partitions {
			//sarama.OffsetNewest：从当前的偏移量开始消费，sarama.OffsetOldest：从最老的偏移量开始消费
			partitionConsumer, err := consumer.ConsumePartition(db.C.Kafka.Topic1, p, sarama.OffsetOldest)
			if err != nil {
				db.L.Errorf("[%s] consumer.ConsumePartition error: %s\n", utils.GetCurrentFunctionName(), err)
				continue
			}
			go func() {
				for m := range partitionConsumer.Messages() {
					fmt.Printf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
				}
			}()
		}
	}
}
