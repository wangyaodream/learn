package main

import (
	"fmt"
    "encoding/json"
    "strings"
)

type Product struct {
    name, category string
    price float64
    // otherNames []string
}

// constructor function
func newProduct(name, category string, price float64) *Product {
    return &Product{name, category, price}
}

func writeName(val struct {
    name, category string
    price float64}) {
        fmt.Println("Name:", val.name)
}

func calcTax(product *Product) {
    if product.price > 100 {
        product.price += product.price * 0.2
    }
}


func main() {

    type StockLevel struct {
        Product
        Alternate Product
        count int
    }

    type Item struct {
        name string
        category string
        price float64
    }


    prod := Product { name: "Kayak", category: "Watersports", price: 275.00}

    array := [1]StockLevel {
        {
            Product: Product{ "Kayak", "Watersports", 275.00},
            Alternate: Product{"Lifejacket", "Watersports", 48.95},
            count: 100,
        },
    }

    fmt.Println("Array:", array[0].Product.name)

    slice := []StockLevel {
        {
            Product: Product{ "Kayak", "Watersports", 275.00},
            Alternate: Product{"Lifejacket", "Watersports", 48.95},
            count: 100,
        },
    }

    fmt.Println("Slice:", slice[0].Product.name)

    kvp := map[string]StockLevel {
        "Kayak": {
            Product: Product{ "kayak", "Watersports", 275.00},
            Alternate: Product{ "Lifejacket", "Watersports", 48.95},
            count: 100,
        },
    }

    fmt.Println("Map:", kvp["Kayak"].Product.name)

    // 在没有给定值的情况会以0值为初始值
    // kayak := Product {
    //     name: "Kayak",
    //     category: "Watersports",
    //     price: 275,
    // }

    // 可以按照顺序来进行数据的初始化
    var builder strings.Builder
    json.NewEncoder(&builder).Encode(struct {
        ProductName string
        ProductPrice float64
    }{
        ProductName: prod.name,
        ProductPrice: prod.price,
    })
    fmt.Println(builder.String())
    
    p1 := Product {
        name: "Kayak",
        category: "Watersports",
        price: 275,
    }

    p2 := &p1

    p1.name = "Original Kayak"

    fmt.Println("P1:", p1.name)
    fmt.Println("P2:", (*p2).name)

    fmt.Println("p1 without tax", p1.price)
    calcTax(&p1)
    fmt.Println("p1 with tax:", p1.price)

    // struct constructor functions
    products := [2]*Product {
        newProduct("Kayak", "Watersports", 275),
        newProduct("Hat", "skiing", 42.50),
    }

    for _, p := range products {
        fmt.Println("Name:", p.name, "Category:", p.category, "Price:", p.price)
    }

}
