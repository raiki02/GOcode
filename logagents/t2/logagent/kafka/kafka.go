package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

var (
	cli    sarama.SyncProducer
	err    error
	ldChan chan *logData
)

type logData struct {
	topic string `json:"topic"`
	data  string `json:"data"`
}

func Init(addr string, chanSize int) {
	ldChan = make(chan *logData, chanSize)

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	cli, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		fmt.Println("kafka init failed, err:", err)
		return
	}
	fmt.Println("kafka init success")
	return
}

func SendToChan(ld *logData) {
	for {
		select {
		case ldChan <- ld:
			sendToKafka(ld)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

func sendToKafka(ld *logData) {
	msg := sarama.ProducerMessage{}
	msg.Topic = ld.topic
	msg.Value = sarama.StringEncoder(ld.data)

	pid, offset, err := cli.SendMessage(&msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("send msg to kafka success, pid:%v offset:%v\n", pid, offset)
}
