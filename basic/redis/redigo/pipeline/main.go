package main

import (
	"fmt"
	"github.com/iralance/go-looklook/basic/redis/redigo/db"
)

func main() {
	c := db.GetRedisPool().Get()
	defer c.Close()

	c.Send("SET", "foo", "bar")
	c.Send("SET", "foo1", "bar1")
	c.Send("GET", "foo")
	c.Send("GET", "foo1")
	c.Flush()
	for i := 0; i < 4; i++ {
		reply, err := c.Receive()
		fmt.Println(reply, err)
	}

}
