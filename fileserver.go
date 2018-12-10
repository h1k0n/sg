package main

import (
    "net/http"
    "fmt"
)

func main() {

    http.Handle("/s/", http.StripPrefix("/s/", http.FileServer(http.Dir("./"))))  

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println(err)
    }
}