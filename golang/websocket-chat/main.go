package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)



var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

type Message struct {
    Username string `json:"username"`
    Message string `json:"message"`
}

var clients = make(map[*websocket.Conn]bool)
var brodcast = make(chan Message)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to the chat Room!")
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer conn.Close()

    clients[conn] = true

    for {
        var msg Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            fmt.Println(err)
            delete(clients, conn)
            return
        }

        brodcast <- msg
    }

}

func handleMessage() {
    for {
        msg := <- brodcast

        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                fmt.Println(err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func main() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/ws", handleConnection)

    go handleMessage()

    fmt.Println("Server started on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic("Error starting server: " + err.Error())
    }
}

