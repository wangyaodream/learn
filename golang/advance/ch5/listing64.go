package main

import (
    "fmt"

    "src/counters"
    "src/entities"
)


func main() {
    // 创建一个未公开类型的变量将其初始化为10
    counter := counters.New(10)

    fmt.Printf("Counter: %d\n", counter)
    fmt.Println(Version)

    u := entities.User{
        Name: "yaoyaoyao",
        email: "yaoyaoyao@email.com",
    }
}

