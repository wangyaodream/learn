package main

import (
    "fmt"
)

func main() {
    names := [3]string {"Kayak", "Lifejacket", "Paddle"}

    otherArray := &names

    names[0] = "Canoe"


    fmt.Println("names:", names)
    fmt.Println("otherArray:", *otherArray)
}

func enumeratArray() {
    names := [3]string {"Kayak", "Lifejacket", "Paddle"}

    for _, value := range names {
        fmt.Println("Value:", value)
    }
}

func createSlice() {
    names := make([]string, 3)
    othernames := []string {"Kayak", "Lifejacket", "Paddle"}

    names[0] = "Kayak"
    names[1] = "Lifejacket"
    names[2] = "Paddle"

    names = append(names, "Hat", "Gloves")

    fmt.Println(names)
    fmt.Println(othernames)

}


