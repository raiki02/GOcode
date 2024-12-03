package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	cli *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

func Init(addr string, timeout time.Duration) (err error) {
	config := clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	}

	cli, err = clientv3.New(config)
	if err != nil {
		fmt.Println("connect etcd failed, err:", err)
		return
	}
	return nil
}

func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Println("get from etcd failed, err:", err)
		return
	}

	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Println("Unmarshal etcd value failed, err:", err)
			return
		}
	}
	return
}
