package redis

import (
	"fmt"
	"time"

	redis1 "github.com/garyburd/redigo/redis"
	"github.com/go-redis/redis"
)

func RedisClient1() {
	//通过 go 向 redis 写入和读取数据
	//1、连接到redis
	conn, err := redis1.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	//4、关闭redis连接
	defer conn.Close()

	//2、通过go 向redis写入数据 string [key-val]
	_, err = conn.Do("Set", "name", "tom哈哈")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	//3、通过go 向redis读取数据 string [key-val]
	r, err := redis1.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	fmt.Println("操作ok", r)
}

func RedisClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // 指定
		Password: "",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	rs, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ping result:", rs)

	rdb.Set("chen_key", "long_value", 5*time.Minute)

	value, err := rdb.Get("chen_key").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("value:", value)
}
