package main

import (
	"fmt"
	"multipkg/lib"
)

func main() {
	fmt.Println("##", lib.NewHello())
	fmt.Println("##", lib.NewHello1())
}
