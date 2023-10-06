package main

import (
    "fmt"
    "strconv"
    "os"
)


func main() {
    var orig string = "ABC"
    var newS string

    fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

    an, err := strconv.Atoi(orig)
    
    check(err)
    fmt.Printf("The integer is %d\n", an)
    an = an + 1
    newS = strconv.Itoa(an)
    fmt.Printf("The new string is : %s\n", newS)
}


func check(e any) {
    if e != nil {
        fmt.Printf("Program stopping with error: %v", e)
        os.Exit(20)
    }
}
