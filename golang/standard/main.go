package main


import (
    "io"
    "os"
    "log"
    _ "fmt"
)

const (
    a = 1 + iota
    b
    c
    d
    e
)

var (
    Trace   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
)

func init() {
    // log.SetPrefix("TRACE:")
    // log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
    file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Failed to open error log file:", err)
    }

    Trace = log.New(io.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
    Info = log.New(io.Discard, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}


func main() {


    // log.Panicln("panic message")

    Trace.Println("I have something standard to say.")
    Info.Println("Specical information")
    Warning.Println("There is something you need to know about")
    Error.Println("Something has failed")

    
}
