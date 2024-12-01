package main

import (
	"bufio"
	"github.com/nsqio/go-nsq"
	"os"
	"strings"
)

var producer *nsq.Producer

func initProducer(str string) error {
	config := nsq.NewConfig()
	producer, _ = nsq.NewProducer(str, config)
	return nil
}

func main() {
	addr := "127.0.0.1:4150"
	initProducer(addr)

	reader := bufio.NewReader(os.Stdin)

	for {
		data, _ := reader.ReadString('\n')
		data = strings.TrimSpace(data)
		producer.Publish("topic_demo", []byte(data))
	}
}
