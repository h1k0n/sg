package main

import (
	"fmt"
	"net/http"
)

var i int

func handler(w http.ResponseWriter, r *http.Request) {
	i++
	fmt.Fprintln(w, "Hey girl, I love you", i)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
