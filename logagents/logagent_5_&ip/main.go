package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"loagent/conf"
	"loagent/etcd"
	"loagent/kafka"
	"loagent/tailf"
	"loagent/utils"
	"sync"
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

	ipStr, err := utils.GetOutBoundIP()
	if err != nil {
		fmt.Println("get ip failed, err:", err)
		return
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)

	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Println("get conf from etcd failed, err:", err)
		return
	}
	fmt.Println("get conf from etcd success, logEntryConf:", logEntryConf)

	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", index, value)
	}
	tailf.Init(logEntryConf)

	newConfChan := tailf.NewConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan)
	wg.Wait()
	//测试put加ip
}
