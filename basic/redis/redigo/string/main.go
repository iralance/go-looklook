package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/iralance/go-looklook/basic/redis/redigo/db"
)

func main() {
	c := db.GetRedisPool().Get()
	defer c.Close()

	_, err := c.Do("MSet", "abc", 100, "efg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}
}
