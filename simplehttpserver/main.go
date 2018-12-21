package main

import (
	"flag"
	"fmt"
	"net/http"
)

var i int

func handler(w http.ResponseWriter, r *http.Request) {
	i++
	fmt.Fprintln(w, "Hey girl, I love you", i)
}

func main() {
	listenAddress := flag.String("addr", ":8090", "listen address")
	flag.Parse()
	http.HandleFunc("/", handler)
	http.ListenAndServe(*listenAddress, nil)
}
