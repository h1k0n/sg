package main

import (
	"fmt"
	"syncpool_collection/collection"
	"time"
)

var buffers = collection.NewBufferPool(65536)

func main() {

	for index := 0; index < 100; index++ {
		go ssse()
		//go ssse()
	}
	time.Sleep(2 * time.Second)
}

func ssse() {
	buf := buffers.Get()
	buf.Write([]byte("hello world"))
	fmt.Println(buf.String())

	buffers.Put(buf)
}
