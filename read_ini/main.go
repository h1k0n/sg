package main

import "fmt"

func main() {
	reader, err := New("conf.cfg")
	if err != nil {
		fmt.Println(err)
	}
	name := reader.Get("# name")
	fmt.Println(name)
}
