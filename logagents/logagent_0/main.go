package main

import (
	"fmt"
	"loagent/kafka"
	"loagent/tailf"
	"time"
)

func run1() {
	for line := range tailf.ReadChan() {
		kafka.SendToKafka("web_log", line.Text)
	}
}

func run() {
	for {
		select {
		case line := <-tailf.ReadChan():
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)

		}
	}
}

func main() {
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Println("init kafka failed, err:", err)
		return
	}
	err = tailf.Init("./my.log")
	if err != nil {
		fmt.Println("init tailf failed, err:", err)
		return
	}

	run()
}
