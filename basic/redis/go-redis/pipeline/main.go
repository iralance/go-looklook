package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/iralance/go-looklook/basic/redis/go-redis/db"
)

var ctx = context.Background()

// wiki https://www.bilibili.com/read/cv15392487
func main() {
	rdb := db.GetRedisClient()
	pipe := rdb.Pipeline()
	t1 := pipe.Get(ctx, "t1")
	fmt.Println("pipe执行前的t1=", t1)
	for i := 0; i < 10; i++ {
		pipe.Set(ctx, fmt.Sprintf("p%d", i), i, 0)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("pipe执行后的t1=", t1)

	cmds, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 10; i++ {
			pipe.Get(ctx, fmt.Sprintf("p%d", i))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(cmds)
	for i, cmd := range cmds {
		fmt.Printf("p%d=%v\n", i, cmd.(*redis.StringCmd).Val())
	}
}
