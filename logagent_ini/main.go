package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"loagent/conf"
	"loagent/kafka"
	"loagent/tailf"
	"time"
)

var (
	cfg = new(conf.AppConf)
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
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)

		}
	}
}

func main() {

	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed, err:", err)
		return
	}

	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Println("init kafka failed, err:", err)
		return
	}
	err = tailf.Init(cfg.TailfConf.FileName)
	if err != nil {
		fmt.Println("init tailf failed, err:", err)
		return
	}

	run()
}
