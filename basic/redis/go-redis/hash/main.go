package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/iralance/go-looklook/basic/redis/go-redis/db"
)

var ctx = context.Background()
var key = "hash1"

type model struct {
	Str1    string   `redis:"str1"`
	Str2    string   `redis:"str2"`
	Int     int      `redis:"int"`
	Bool    bool     `redis:"bool"`
	Ignored struct{} `redis:"-"`
}

// wiki https://www.bilibili.com/read/cv15392487
// wiki https://redis.uptrace.dev/guide/scanning-hash-fields.html
func main() {
	rdb := db.GetRedisClient()

	if _, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "str1", "hello")
		rdb.HSet(ctx, key, "str2", "world")
		rdb.HSet(ctx, key, "int", 123)
		rdb.HSet(ctx, key, "bool", 1)
		return nil
	}); err != nil {
		panic(err)
	}

	var model1 model
	// Scan all fields into the model.
	if err := rdb.HGetAll(ctx, key).Scan(&model1); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", model1)

	var model2 model
	// Scan a subset of the fields.
	if err := rdb.HMGet(ctx, key, "str1", "int").Scan(&model2); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", model2)

}
