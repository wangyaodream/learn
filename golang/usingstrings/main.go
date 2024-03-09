package main

import "fmt"

func Printfln(template string, values ...interface{}) {
    fmt.Printf(template + "\n", values...)
}

func getProductName(index int) (name string, err error) {
    if (len(Products) > index) {
        name = fmt.Sprintf("Name of product: %v", Products[index].Name)
    } else {
        err = fmt.Errorf("Error for index %v", index)
    }
    return
}

func main() {
    // name, _ := getProductName(1)
    // fmt.Println(name)
    //
    // _, err := getProductName(10)
    // fmt.Println(err.Error())

    // 输出值, 如果方法实现了String()接口，则以自定义的方式显示内容
    Printfln("Value: %v", Kayak)
    // 输出变量原本的内容
    Printfln("Go syntax: %#v", Kayak)
    // 输出类型
    Printfln("Type: %T", Kayak)

    // 输出值的结构
    Printfln("Value with fields: %+v", Kayak)

    number := 279.00

    Printfln("Binary: %b", number) 
    Printfln("Decimal: %d", number) 
    Printfln("Octal: %o, %O", number, number) 
    Printfln("Hexadecimal: %x, %X", number, number)

    Printfln("Decimal without exponent: >>%8.2f<<", number)
    Printfln("Decimal without exponent: >>%.2f<<", number)

    Printfln("Sign: >>%+.2f<<", number)
    Printfln("Zeros for padding: >>%010.2f<<", number)
    Printfln("Right padding: >>%-8.2f<<", number)
    // name := "Kayak"
    // Printfln("Stiring: %s", name)
    // Printfln("Character: %c", []rune(name)[1])
    // Printfln("Unicode: %U", []rune(name)[0])
    //
    // Printfln("Bool: %t", len(name) > 1)
    // Printfln("Bool: %t", len(name) > 100)
    // 
    // Printfln("Pointer: %p", &name)

    var name string
    var category string
    var price float64

    source := "Lifejacket Watersports 48.95"
    n, err := fmt.Sscan(source, &name, &category, &price)

    if (err == nil) {
        Printfln("Scanned %v values", n)
        Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
    } else {
        Printfln("Error: %v", err.Error())
    }


}
