package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)


func startTestHTTPServer() *httptest.Server {
    ts := httptest.NewServer(
        http.HandlerFunc(
            func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "Hello world")
            }))
    return ts
}

func TestFetchRemoteResource(t *testing.T) {
    ts := startTestHTTPServer()
    defer ts.Close()

    expected := "Hello world"
    client := createHTTPClientWithTimeout(10 * time.Second)
    data, err := fetchRemoteResource(client, ts.URL)
    if err != nil {
        t.Fatal(err)
    }
    if expected != string(data) {
        t.Errorf("Expected to be: %s, Got: %s", expected, data)
    }
}
