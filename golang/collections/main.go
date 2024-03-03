package main

import (
    "fmt"
    "sort"
    "reflect"
)

func main() {
    names := [3]string {"Kayak", "Lifejacket", "Paddle"}

    // otherArray := &names

    names[0] = "Canoe"


    // fmt.Println("names:", names)
    // fmt.Println("otherArray:", *otherArray)

    // useCopy()
    // sortSlice()
    // compare()
    workWithMap()
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

func allocating() {
    // 3是length, 6是capacity
    names := make([]int, 3, 6)

    names[2] = 10
    // names[3] = 100 这样会出错，超过了lenght
    names = append(names, 99, 100, 1000, 5000)

    fmt.Println(names)

    fmt.Println("len:", len(names))
    fmt.Println("cap:", cap(names))

}

func useCopy() {
    products := [4]string {"Kayak", "Lifejacket", "Paddle", "Hat"}

    allNames := products[1:]
    // someNames必须进行初始化，否则无法进行copy操作
    someNames := make([]string, 2)
    copy(someNames, allNames)

    fmt.Println("someNames:", someNames)
    fmt.Println("allNames:", allNames)
}

func sortSlice() {
    products := []string {"Kayak", "Lifejacket", "Paddle", "Hat"}
    
    sort.Strings(products)

    for index, value := range products {
        fmt.Println("Index:", index, "Value:", value)
    }
}

func compare() {
    p1 := []string {"Kayak", "Lifejacket", "Paddle", "Hat"}
    p2 := p1

    fmt.Println("Equal:", reflect.DeepEqual(p1, p2))
}

func workWithMap() {
    // products := make(map[string]float64, 10)
    // 
    // products["Kayak"] = 279
    // products["Lifejacket"] = 48.95

    products := map[string]float64 {
        "Kayak": 279,
        "Lifejacket": 48.95,
        "Hat": 0
    }

    value, ok := products["HAt"]

    if (ok) {
        fmt.Println("Stored value:", value)
    } else {
        fmt.Println("No stored value")
    }

    fmt.Println("Map size:", len(products))
    fmt.Println("Price:", products["Kayak"])
    fmt.Println("Price:", products["Hat"])
}
