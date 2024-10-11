package main

import (
	"fmt"
	"time"
)


func main() {
    // create a channel without buffer
    messages := make(chan string)

    go func() {
        for i := 0; i < 5; i++ {
            time.Sleep(1 * time.Second)
            messages <- fmt.Sprintf("Message-%d", i)
        }

        close(messages)
    }()

    for msg := range messages {
        fmt.Println("Received: ", msg)
    }

    fmt.Println("Done!")
}
