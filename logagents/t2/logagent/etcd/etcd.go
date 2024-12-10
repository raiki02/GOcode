package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type LogEntry struct {
	Path  string
	Topic string
}

var (
	cli *clientv3.Client
	err error
)

func Init(addr string) {
	config := clientv3.Config{}
	config.Endpoints = []string{addr}

	cli, err = clientv3.New(config)
	if err != nil {
		fmt.Println("etcd init failed, err:", err)
		return
	}
	fmt.Println("etcd init success")
	return
}

func GetConfs(key string) (LogEntrys []*LogEntry) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	res, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Println("get conf from etcd failed, err:", err)
		return
	}

	for _, kvs := range res.Kvs {
		json.Unmarshal(kvs.Value, &LogEntrys)
	}
	return
}
