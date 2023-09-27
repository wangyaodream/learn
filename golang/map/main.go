package main


import "fmt"


func main() {
    // 创建一个映射
    // dict := make(map[string]int)

    // 创建两个键值对初始化映射
    // d := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

    // 通过声明映射创建一个nil映射
    // fmt.Println(exists)
    colors := map[string]string {
        "AliceBlue":    "#f0f8ff",
        "Coral":        "#ff7F50",
        "DarkGray":     "#a9a9a9",
        "ForestGreen":  "#228b22",
    }

    for _, value := range colors {
        fmt.Println(value)
    }


    fmt.Println("End!")
}
