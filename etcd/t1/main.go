package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	defer cli.Close()

	fmt.Println("connect to etcd success")

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, _ = cli.Put(ctx, "/demo/demo1_key", "demo1_value")
	cancel()

	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, _ := cli.Get(ctx, "/demo/demo1_key")
	//resp, _ := cli.Get(ctx, "/demo/", clientv3.WithPrefix())// get all keys
	cancel()

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	//edit
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	_, _ = cli.Put(ctx, "/demo/demo1_key", "demo1_value_new", clientv3.WithPrevKV())

	//delete
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	_, _ = cli.Delete(ctx, "/demo/demo1_key")
}
