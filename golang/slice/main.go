package main

import "fmt"


func main() {
    fmt.Println("Start:")
    slice := []int{10,20,30,40}

    for index, value := range slice {
        fmt.Printf("Value: %d Value-addr: %x ElemAddr: %x\n", 
        value, &value, &slice[index])
    }

    // traditional for loop
    fmt.Println("Traditional:")
    for index := 2; index < len(slice); index++ {
        fmt.Printf("Index %d Value: %d\n", index, slice[index])
    }

    slice_b := [][]int{{10}, {100, 200}}
    fmt.Printf("slice_b: %v", slice_b)
}
