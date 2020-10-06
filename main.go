package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("received request")
        fmt.Fprintf(w, "Hello Golang Docker!!")
    })
    
    log.Println("Start server")
    server:= &http.Server {
        Addr: ":8000",
    }
    if err:= server.ListenAndServe(); err != nil {
        log.Println(err)
    }
}