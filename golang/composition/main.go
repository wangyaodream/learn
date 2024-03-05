package main

import (
    "fmt"
    "composition/store"
)

func main() {

    // rentals := []*store.RentalBoat {
    //     store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
    //     store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
    //     store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
    // }
    //
    // for _, r := range rentals {
    //     fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2), "Captain:", r.Captain)
    // }
    //
    // boats := []*store.Boat {
    //     store.NewBoat("Kayak", 275, 1, false),
    //     store.NewBoat("Canoe", 400, 3, false),
    //     store.NewBoat("Tender", 650.25, 2, true),
    // }
    // 
    // kayak := store.NewProduct("Kayak", "Watersports", 275)
    // lifejacket := &store.Product{Name: "lifejacket", Category: "Watersports"}
    //
    // for _, p := range []*store.Product {kayak, lifejacket} {
    //     fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
    // }
    //
    // for _, b := range boats {
    //     // fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Name)
    //     fmt.Println("Boat:", b.Name, "Price:", b.Price(0.2))
    // }

    // product := store.NewProduct("kayak", "Watersports", 279)
    //
    // deal := store.NewSpecialDeal("Weekend Sepcial", product, 50)
    //
    // Name, price, Price := deal.GetDetails()
    //
    // fmt.Println("Name:", Name)
    // fmt.Println("Price field:", price)
    // fmt.Println("Price method:", Price)

    // 使用interface来接纳不同类型的数据
    products := map[string]store.ItemForSale {
        "Kayak": store.NewBoat("Kayak", 279, 1, false),
        "Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
    }

    for key, p := range products {
        // fmt.Println("Key:", key, "Price:", p.Price(0.2))
        switch item := p.(type) {
            // case *store.Product:
            //     fmt.Println("Name:", item.Name, "Category:", item.Category, "Price:", item.Price(0.2))
            // case *store.Boat:
            //     fmt.Println("Name:", item.Name, "Category:", item.Category, "Price:", item.Price(0.2))
            case store.Describable:
                fmt.Println("Name:", item.GetName(), "Category:", item.GetCategory(), "Price:", item.(store.ItemForSale).Price(0.2))
            default:
                fmt.Println("Key:", key, "Price:", p.Price(0.2))
        }
    }

}
