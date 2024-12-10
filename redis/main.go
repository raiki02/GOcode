package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

/*
func initRedis_old() {
	c, err := redis.Dail("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis connect error: ", err)
	}

	fmt.Println("redis connect success")

	defer c.Close()
}
*/

func initRedis() {

}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	err := rdb.Set("key1", "value1", 0).Err()
	if err != nil {
		fmt.Println("redis set error: ", err)
	}

	val, err := rdb.Get("key1").Result()
	if err != nil {
		fmt.Println("redis get error: ", err)
	}
	fmt.Println("key1: ", val)

	err = rdb.Get("key2").Err()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		fmt.Println("redis get error: ", err)
	} else {
		fmt.Println("key2: ", val)
	}

}
