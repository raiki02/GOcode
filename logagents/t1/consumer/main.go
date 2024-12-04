package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

func main() {
	consumer, _ := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	defer consumer.Close()

	partitionList, _ := consumer.Partitions("web_log")
	fmt.Println("partitionList : ", partitionList)
	
	for partition := range partitionList {
		pc, _ := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		defer pc.AsyncClose()

		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition: %d Offset: %d Key: %s Value: %s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
}
