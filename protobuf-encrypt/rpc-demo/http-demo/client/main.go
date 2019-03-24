package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Params struct {
	Width  int
	Height int
}

func main() {
	rpc, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		ret := 0
		err = rpc.Call("Rect.Area", Params{50, 100}, &ret)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ret)

		time.Sleep(time.Second)
	}
}
