package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/iralance/go-looklook/basic/redis/go-redis/db"
	"time"
)

var ctx = context.Background()

// wiki https://www.bilibili.com/read/cv15392487
// diff go-redis vs redigo  https://redis.uptrace.dev/guide/go-redis-vs-redigo.html
func main() {
	rdb := db.GetRedisClient()
	// set key value
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// get key
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	cmd := rdb.Do(ctx, "set", "key3", "value3", "ex", 10, "nx")
	res, err := cmd.Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("do失败")
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("res=", res.(string))
	}
	fmt.Println(res, err)
	//其实对redis true = 1
	rdb.Do(ctx, "set", "b1", true)
	rdb.Do(ctx, "set", "b2", 0)
	b, err := rdb.Do(ctx, "get", "b1").Bool()
	if err == nil {
		fmt.Println("get b1=", b)
	} else {
		fmt.Println("get b1 err", err)
	}

	boolSlice, err := rdb.Do(ctx, "mget", "b1", "b2").BoolSlice()
	if err == nil {
		fmt.Println("mget b1 b2=", boolSlice)
	}
	rdb.Set(ctx, "t1", time.Now(), 0)
	t, err := rdb.Get(ctx, "t1").Time()
	fmt.Println(t, err)
}
