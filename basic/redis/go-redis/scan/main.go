package main

import (
	"context"
	"fmt"
	"github.com/iralance/go-looklook/basic/redis/go-redis/db"
)

var ctx = context.Background()

// wiki https://www.bilibili.com/read/cv15392487
// diff go-redis vs redigo  https://redis.uptrace.dev/guide/go-redis-vs-redigo.html
func main() {
	rdb := db.GetRedisClient()

	iter := rdb.Scan(ctx, 0, "p*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Printf("key=%v, value=%v\n", iter.Val(), rdb.Get(ctx, iter.Val()).String())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}
