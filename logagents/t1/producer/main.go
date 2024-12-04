package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/hpcloud/tail"
	"time"
)

var (
	tails  *tail.Tail
	client sarama.SyncProducer
	line   *tail.Line
)

func InitTailFile() {
	fileName := "./t1.log"
	config1 := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Poll:      true,
		MustExist: false,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
	}

	tails, _ = tail.TailFile(fileName, config1)

}

func run() {

}

func InitProducer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, _ = sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
}

func SendToKafka(text string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder(text)

	pid, offset, _ := client.SendMessage(msg)

	fmt.Printf("pid : %v , offset  : %v \n", pid, offset)
}

func main() {
	InitTailFile()
	fmt.Println("init tail file success")

	InitProducer()
	fmt.Println("init kafka success")

	for {
		select {
		case line = <-tails.Lines:
			fmt.Println("line.Text : ", line.Text)
			SendToKafka(line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}
