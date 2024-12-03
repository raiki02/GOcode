package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"loagent/conf"
	"loagent/etcd"
	"loagent/kafka"
	"loagent/tailf"
	"time"
)

var (
	cfg = new(conf.AppConf)
)

func main() {

	err := ini.MapTo(cfg, "./conf/conf.ini")
	if err != nil {
		fmt.Println("load ini failed, err:", err)
		return
	}
	fmt.Println("cfg map yes")
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.MaxSize)
	if err != nil {
		fmt.Println("init kafka failed, err:", err)
		return
	}
	fmt.Println("kafka init yes")
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Println("init etcd failed, err:", err)
		return
	}

	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Println("get conf from etcd failed, err:", err)
		return
	}

	fmt.Println("get conf from etcd success, logEntryConf:", logEntryConf)

	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", index, value)
	}
	tailf.Init(logEntryConf)

}
