package main

import (
    "log"
    "os"
    "net/http"

    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

func mainLogic(w http.ResponseWriter, r *http.Request) {
    log.Println("Processing request!")
    w.Write([]byte("OK"))

    log.Println("Finished processing request")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", mainLogic)
    loggedRouter := handlers.LoggingHandler(os.Stdout, r)
    http.ListenAndServe(":8000", loggedRouter)
}
