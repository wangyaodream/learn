package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{}

func socketHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("Error during connection upgradation:", err)
        return
    }
    defer conn.Close()

    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error during message reading", err)
        }

        log.Printf("Received: %s", message)
        err = conn.WriteMessage(messageType, message)
        if err != nil {
            log.Println("Error during message writing:", err)
            break
        }
    }
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Index Page")
}

func main() {
    http.HandleFunc("/socket", socketHandler)
    http.HandleFunc("/", home)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
