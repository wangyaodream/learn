package main


import (
    "io"
    "io/ioutil"
    "os"
    "log"
    "fmt"
)

const (
    a = 1 + iota
    b
    c
    d
    e
)

func init() {
    log.SetPrefix("TRACE:")
    log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}


func main() {
    log.Println("message")

    log.Fatalln("fatal message")

    // log.Panicln("panic message")

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(e)
    
}
