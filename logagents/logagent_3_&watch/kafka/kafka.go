package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

var (
	client      sarama.SyncProducer
	logDataChan chan *logData
)

type logData struct {
	topic string
	data  string
}

func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		return
	}

	logDataChan = make(chan *logData, maxSize)
	go sendToKafka()
	return
}

func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)

			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				return
			}
			fmt.Printf("pid : %v , offset : %v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}

	}
}

// ---------------------------------------------------
func SendToKafka(topic, data string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		return
	}
	fmt.Printf("pid : %v , offset : %v\n", pid, offset)
}
