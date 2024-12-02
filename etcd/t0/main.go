package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect to etcd failed, err:", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"c:/tmp/nginx.log","topic":"web_log"},{"path":"d:/xxx/redis.log","topic":"redis_log"}]`
	_, err = cli.Put(ctx, "/logagent_ini_etcd/collect_config", value)
	if err != nil {
		fmt.Println("put to etcd failed, err:", err)
		return
	}
	cancel()

	//ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	//res, err := cli.Get(ctx, "baodelu")
	//
	//cancel()
	//if err != nil {
	//	fmt.Println("get from etcd failed, err:", err)
	//	return
	//}
	//
	//for _, ev := range res.Kvs {
	//	fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	//}

	//watchchan := cli.Watch(context.Background(), "baodelu")
	//for wres := range watchchan {
	//	for _, ev := range wres.Events {
	//		fmt.Printf("Type: %s Key: %s Value: %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	//	}
	//}
}
