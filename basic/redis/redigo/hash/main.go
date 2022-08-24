package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/iralance/go-looklook/basic/redis/redigo/db"
)

type model struct {
	Str1    string   `redis:"str1"`
	Str2    string   `redis:"str2"`
	Int     int      `redis:"int"`
	Bool    bool     `redis:"bool"`
	Ignored struct{} `redis:"-"`
}

//https://pkg.go.dev/github.com/gomodule/redigo/redis#pkg-examples
func main() {
	c := db.GetRedisPool().Get()
	defer c.Close()

	model1 := model{
		Str1: "张三",
		Str2: "李四",
		Int:  20,
		Bool: false,
	}
	model2 := model{}
	if _, err := c.Do("HMSET", redis.Args{}.Add("model1").AddFlat(&model1)...); err != nil {
		fmt.Println(err)
		return
	}

	m := map[string]string{
		"title":  "Example2",
		"author": "Steve",
		"body":   "Map",
	}

	if _, err := c.Do("HMSET", redis.Args{}.Add("id2").AddFlat(m)...); err != nil {
		fmt.Println(err)
		return
	}

	v, err := redis.Values(c.Do("HGETALL", "model1"))
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := redis.ScanStruct(v, &model2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", model2)

}
