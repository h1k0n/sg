package main

import (
	"crypto/rc4"
	"fmt"
	"log"
)

func main() {
	key := []byte("dsadsad")
	c, err := rc4.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
	}
	src := []byte("asdsad")
	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)
	fmt.Println(string(src), string(dst))
}
