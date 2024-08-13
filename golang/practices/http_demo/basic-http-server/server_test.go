package main

import (
	"net/http"
	"testing"
)


func TestServer(t *testing.T) {
    tests := []struct{
        name string
        path string
        expected string
    } {
        {
            name: "index",
            path: "/api",
            expected: "Hello world",
        },
        {
            name: "healthcheck",
            path: "/health",
            expected: "ok",
        },
    }

    mux := http.NewServeMux()
    
}
