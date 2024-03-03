package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Test(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "This is a test")
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category is: %v\n", vars["category"])
    fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
    queryParams := r.URL.Query()
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Got parameter id:%s!\n", queryParams["id"][0])
    fmt.Fprintf(w, "Got parameter category:%s!\n", queryParams["category"][0])
}

func main() {
    // create a new router
    r := mux.NewRouter()
    // r.Host("aaa.bbb.ccc")
    r.HandleFunc("/articles/", QueryHandler)
    r.Queries("id", "category")

    r.StrictSlash(true)
    r.HandleFunc("/test", Test)
    // s := r.PathPrefix("/articles").Subrouter()
    // s.HandleFunc("{id}/settings", settingsHandler)
    // s.HandleFunc("{id}/detail", detailsHandler)
    // Attach a path with handler
    // r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
    srv := &http.Server{
        Handler: r,
        Addr: "127.0.0.1:8000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
    }
    log.Fatal(srv.ListenAndServe())
}
