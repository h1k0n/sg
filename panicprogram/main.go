package main

import (
	"time"
)

func main() {
	for {
		time.Sleep(10 * time.Second)
		panic("error")
	}
}
