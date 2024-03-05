package main

import (
    "fmt"
    "packages/store"
    . "packages/fmt"
    "packages/store/cart"
    _ "packages/data"
    "github.com/fatih/color"
)

func main() {
    
    product := store.NewProduct("Kayak", "Watersports", 279)

    cart := cart.Cart {
        CustomerName: "Alice",
        Products: []store.Product{ *product },        
    }

    // fmt.Println("Name:", product.Name)
    // fmt.Println("Category:", product.Category)
    // // 通过定义的方法来访问price的值，并不能直接读取price本身
    // fmt.Println("Price:", product.Price())
    // fmt.Println("currencyFmt:", ToCurrency(product.Price()))

    color.Green("Name:" + cart.CustomerName)
    color.Cyan("Total:" + ToCurrency(cart.GetTotal()))

    fmt.Println("Name:", cart.CustomerName)
    fmt.Println("Total:", ToCurrency(cart.GetTotal()))
}
