package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)


func process(conn net.Conn) {
    defer conn.Close()
    for {
        reader := bufio.NewReader(conn)
        var buf [128]byte
        n, err := reader.Read(buf[:])
        if err != nil {
            log.Println("read from client failed, err:", err)
            break
        }

        recvStr := string(buf[:n])
        fmt.Println("receive data from client:", recvStr)
        conn.Write([]byte(recvStr))
    }
}

func main() {
    listen, err := net.Listen("tcp", "127.0.0.1:20000")
    if err != nil {
        log.Println("listen failed, err:", err)
        return
    }
    log.Println("listening...")
    for {
        conn, err := listen.Accept()
        if err != nil {
            fmt.Println("accept failed, err:", err)
            continue
        }

        go process(conn)
    }

}
