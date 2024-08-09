package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)



func startTestPackgeServer() *httptest.Server {
    pkgData := `[
    {"name": "package1", "version": "1.1"},
    {"name": "package2", "version": "3.1"}
    ]`

    ts := httptest.NewServer(
        http.HandlerFunc(
            func(w http.ResponseWriter,r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprint(w, pkgData)
            }))
    return ts
}

func TestFetchPakcageData(t *testing.T) {
    ts := startTestPackgeServer()
    defer ts.Close()
    packages, err := fetchPackageData(ts.URL)
    if err != nil {
        t.Fatal(err)
    }

    if len(packages) != 2 {
        t.Fatalf("Excepted 2 packages, Got back: %d", len(packages))
    }
}
