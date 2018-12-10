package main

import (
	"fmt"
	"sync"
	"time"
)

// 一个[]byte的对象池，每个对象为一个[]byte
var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	a := time.Now().Unix()
	// 不使用对象池
	var v *[]byte
	for i := 0; i < 1000000000; i++ {
		obj := make([]byte, 1024)
		v = &obj
	}
	b := time.Now().Unix()
	// 使用对象池
	for i := 0; i < 1000000000; i++ {
		obj := bytePool.Get().(*[]byte)
		v = obj
		bytePool.Put(obj)
	}
	c := time.Now().Unix()
	fmt.Println(v)
	fmt.Println("without pool ", b-a, "s")
	fmt.Println("with    pool ", c-b, "s")
}

// without pool  224 s
// with    pool  18 s
