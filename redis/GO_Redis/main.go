package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func InitRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err = redisdb.Ping().Result()
	return
}

func setScore() {
	err := redisdb.Set("score", 100, 0).Err()
	if err != nil {
		panic(err)
	}
}

func getScore() {
	val, err := redisdb.Get("score").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("score is ", val)

	val2, err := redisdb.Get("name").Result()
	if err != nil {
		fmt.Println("get name err ", err)
		return
	} else if err == redis.Nil {
		fmt.Println("name not exist ")
		return
	} else {
		fmt.Println("name = ", val2)
	}
}

func redisZset() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 9.0, Member: "js"},
		redis.Z{Score: 20.0, Member: "En"},
		redis.Z{Score: 05.0, Member: "Ch"},
		redis.Z{Score: 60.0, Member: "Fr"},
		redis.Z{Score: 920.0, Member: "Ru"},
		redis.Z{Score: 2620.0, Member: "Jp"},
		redis.Z{Score: 9270.0, Member: "cpp"},
	}

	// ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)

	// ret := redisdb.ZRevRangeWithScores(ctx, 0, 7).Val()
	// for _, v := range ret {
	// 	fmt.Println(v.Member, v.Score)
	// }
	num, err := redisdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Println("zAdd err ", err)
		return
	}
	fmt.Println("zAdd yes, added", num)
	newScore, err := redisdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Println("zIncrby err ", err)
		return
	}
	fmt.Println("golang score is now", newScore)

}

func main() {
	err := InitRedis()
	if err != nil {
		fmt.Println("init redis err: ", err)
		return
	}
	fmt.Println("init redis yes")

	setScore()
	getScore()
	redisZset()
}
