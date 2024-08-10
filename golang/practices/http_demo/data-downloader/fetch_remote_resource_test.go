package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
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
    data, err := fetchRemoteResource(ts.URL)
    if err != nil {
        t.Fatal(err)
    }
    if expected != string(data) {
        t.Errorf("Expected to be: %s, Got: %s", expected, data)
    }
}
