package main

import "fmt"


func main() {
    arrayA := [2]int {100, 200}
    var arrayB [2]int

    arrayB = arrayA
    
    fmt.Printf("arrayA: %p, %v\n", &arrayA, arrayA)
    fmt.Printf("arrayB: %p, %v\n", &arrayB, arrayB)

    testArray(arrayA)
    // 会打印出来三个不同的地址，表示分配了三次内存空间

    fmt.Println(arrayA)
    testArray2(&arrayA)
    fmt.Println(arrayA)
}

func testArray(x [2]int) {
    fmt.Printf("func Array: %p, %v\n", &x, x)
    x[0] = 999
}

func testArray2(x *[2]int) {
    fmt.Printf("func Array 2: %p, %v\n", &x, x)
    (*x)[0] = 999 
}
