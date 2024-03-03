package main

import (
    "log"
    "net"
    "net/http"
    "net/rpc"
    // "time"
)

type Args struct {}
type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
    // Fill reply pointer to send the data back
    // *reply = time.Now().Unix()
    *reply = 999
    return nil
}

func main() {
    // Create a new RPC server
    timeserver := new(TimeServer)
    rpc.Register(timeserver)
    rpc.HandleHTTP()
    // Listen for requests on port 1234
    l, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    http.Serve(l, nil)
}
