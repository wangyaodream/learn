package main

import "database/sql"

type Category struct {
    Id int
    Name string
}

type Product struct {
    Id int
    Name string
    Category
    Price float64
}

func queryDatabase(db *sql.DB) []Product {
    products := []Product {}
    rows, err := db.Query(`
        SELECT Products.Id, Products.Name, Products.Price,
            Categories.Id as Cat_Id, Categories.Name as CatName
            FROM Products, Categories
        WHERE Products.Category = Categories.Id`)
    if (err == nil) {
        for (rows.Next()) {
            p := Product{}
            scanErr := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
            if (scanErr == nil) {
                products = append(products, p)
            } else {
                Printfln("Scan error: %v", scanErr)
                break
            }
        }
    } else {
        Printfln("Error: %v", err)
    }
    return products
}

func main() {
    // listDrivers()
    db, err := openDatabase()
    if (err == nil) {
        products := queryDatabase(db)
        for i, p := range products {
            Printfln("#%v: %v", i, p)
        }
        db.Close()
    } else {
        panic(err)
    }
}
