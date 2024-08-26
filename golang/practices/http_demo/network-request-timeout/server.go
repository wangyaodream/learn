package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)


func handlePing(w http.ResponseWriter, h *http.Request) {
    log.Println("ping: Got a request")
    fmt.Fprint(w, "Pong")
}

func doSomeWork() {
    time.Sleep(2 * time.Second)
}

func handleUserApi(w http.ResponseWriter, r *http.Request) {
    log.Println("I started processing the request")

    doSomeWork()

    req, err := http.NewRequestWithContext(
        r.Context(),
        "GET",
        "http://localhost:5000/ping", nil,
    )
    if err != nil {
        http.Error(
            w, err.Error(),
            http.StatusInternalServerError,
        )
    }

    client := &http.Client{}
    log.Println("Outgoing HTTP request")

    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Error making request: %v\n", err)
        http.Error(
            w, err.Error(),
            http.StatusInternalServerError,
        )
    }

    defer resp.Body.Close()
    data, _ := io.ReadAll(resp.Body)
    fmt.Fprintln(w, string(data))
    log.Println("I finished processing the request")
}

func main() {
    listenAddr := os.Getenv("LISTEN_ADDR")
    if len(listenAddr) == 0 {
        listenAddr = ":5000"
    }

    timeoutDuration := 1 * time.Second

    userHandler := http.HandlerFunc(handleUserApi)
    hTimeout := http.TimeoutHandler(
        userHandler,
        timeoutDuration,
        "I ran out of time",
    )
    mux := http.NewServeMux()
    mux.Handle("/api/users/", hTimeout)
    mux.HandleFunc("/ping", handlePing)

    log.Fatal(http.ListenAndServe(listenAddr, mux))
}
