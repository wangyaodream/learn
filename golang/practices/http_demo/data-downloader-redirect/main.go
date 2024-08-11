package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)


func fetchRemoteResource(client *http.Client, url string) ([]byte, error) {
    r, err := client.Get(url)
    if err != nil {
        return nil, err
    }
    defer r.Body.Close()

    return io.ReadAll(r.Body)
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
    if len(via) >= 1 {
        return errors.New("stopped after 1 redirect")
    }
    return nil
}

func createHTTPClientWithTimeout(d time.Duration) *http.Client {
    client := http.Client{Timeout: d, CheckRedirect: redirectPolicyFunc}
    return &client
}

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stdout, "Must specify a HTTP URL to get data from")
        os.Exit(1)
    }

    client := createHTTPClientWithTimeout(time.Second * 15)
    body, err := fetchRemoteResource(client, os.Args[1])
    if err != nil {
        fmt.Fprintf(os.Stdout, "%v\n", err)
        os.Exit(1)
    }

    fmt.Fprintf(os.Stdout, "%s\n", body)
}
