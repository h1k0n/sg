//https://play.golang.org/p/f-VQX8Lcb09
package main

import (
	"fmt"
)

func main() {
	var n Noun
	a := A{
		name: "a",
	}
	b := B{
		name: "b",
	}
	fmt.Println("Hello, playground")
	n = returnType(a, b, "b")
	//n = a
	fmt.Println(n.Name())
}

type Noun interface {
	Name() string
}

type A struct {
	name string
}

type B struct {
	name string
}

func (a A) Name() string {
	return a.name
}

func (b B) Name() string {
	return b.name
}
func returnType(a, b interface{}, name string) Noun {
	aa, ok := a.(A)
	if ok {
	}
	bb, ok := b.(B)
	if ok {
	}
	//fmt.Println(a, ok)
	if name == "a" {
		return aa
	}

	return bb
}
