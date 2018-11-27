package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func subscribe() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	psc := redis.PubSubConn{c}
	psc.Subscribe("redChatRoom")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}
func main() {
	go subscribe()
	go subscribe()
	go subscribe()
	go subscribe()
	go subscribe()

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	for {
		var s string
		fmt.Scanln(&s)
		_, err := c.Do("PUBLISH", "redChatRoom", s)
		if err != nil {
			fmt.Println("pub err: ", err)
			return
		}
	}
}
