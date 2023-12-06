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

    a := entities.Admin{
        Rights: 2,
    }

    // 设置未公开的user中的值
    a.Name = "wangyao"
    a.Email = "wangyao@email.com"
}

