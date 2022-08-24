package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/iralance/go-looklook/basic/redis/redigo/db"
)

func main() {
	c := db.GetRedisPool().Get()
	defer c.Close()

	_, err := c.Do("lpush", "book_list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
}
