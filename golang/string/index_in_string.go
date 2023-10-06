package main

import (
    "fmt"
    "strings"
)


func main() {
    var str string = "Hi, I'm Marc, Hi."

    fmt.Printf("The position of \"Marc\"")
    fmt.Printf("%d\n", strings.Index(str, "Marc"))
    fmt.Println(str)

    new_str := strings.Replace(str, "Hi", "Hello", 1)
    fmt.Println(new_str)

    manyG := "ggggggggggg"

    fmt.Printf("%d\n", strings.Count(manyG, "g"))
}
