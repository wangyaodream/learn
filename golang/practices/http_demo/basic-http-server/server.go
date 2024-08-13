package main

import (
	"log"
	"net/http"
)


func main() {
    listenAddr := "localhost:5000"
    if len(listenAddr) == 0 {
        listenAddr = ":8080"
    }
    log.Fatal(http.ListenAndServe(listenAddr, nil))
}
