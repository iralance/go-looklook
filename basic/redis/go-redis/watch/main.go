package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/iralance/go-looklook/basic/redis/go-redis/db"
)

var ctx = context.Background()

// wiki https://www.bilibili.com/read/cv15392487
// diff go-redis vs redigo  https://redis.uptrace.dev/guide/go-redis-vs-redigo.html
func main() {
	rdb := db.GetRedisClient()

	for i := 0; i < 10; i++ {
		err := rdb.Watch(ctx, func(tx *redis.Tx) (err error) {
			pipeline := rdb.Pipeline()
			err = pipeline.IncrBy(ctx, "p1", 100).Err()
			if err != nil {
				return err
			}
			err = pipeline.DecrBy(ctx, "p1", 100).Err()
			if err != nil {
				return
			}
			_, err = pipeline.Exec(ctx)
			return
		})

		if err == nil {
			fmt.Println("事务commit成功")
			break
		} else if err == redis.TxFailedErr {
			fmt.Println("事务执行失败，这次是第", i, "次执行")
			continue
		} else {
			panic(err)
		}
	}
}
