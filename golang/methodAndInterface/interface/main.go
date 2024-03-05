package main

import (
    "fmt"
)

type Expense interface {
    getName() string
    getCost(annual bool) float64
}

type Foo struct {
    name string
    price float64
}

func (f Foo) getName() string {
    f.name = "this is foo"
    return f.name
}

func (f Foo) getCost(_ bool) float64 {
    return f.price
}

func calcTotal(expenses []Expense) (total float64) {
    for _, item := range expenses {
        total += item.getCost(true)
    }
    return
}

type Account struct { 
    accountNumber int
    expenses []Expense
}

func main() {

    // product := Product { "Kayak", "Watersports", 292}

    // var expense Expense = product
    // var expense Expense = &product
    //
    // product.price = 100
    //
    // // expense会复制product的数据，所以在调用getCost方法时并不会使用改变后的数据
    // fmt.Println("Product field value:", product.price)
    // fmt.Println("Expense method result:", expense.getCost(false))

    // account := Account {
    //     accountNumber: 12345,
    //     expenses: []Expense {
    //         &Product{"Kayak", "Watersports", 300},
    //         Service{"Boat Cover", 12, 100},
    //     },
    // }
    //
    // expenses := []Expense {
    //     &Product{"kayak", "Watersports", 275},
    //     Service{"Boat Cover", 12, 89.50},
    // }
    //
    // for _, expense := range expenses {
    //     fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
    // }
    //
    // fmt.Println("Total:", calcTotal(expenses))
    //
    // for _, expense := range account.expenses {
    //     fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
    // }
    // 
    // fmt.Println("Total:", calcTotal(account.expenses))

    // kayak := Product {"Kayak", "Watersports", 275}
    // insurance := Service {"Boat Cover", 12, 89.50}
    //
    // fmt.Println("Product:", kayak.name, "Price:", kayak.price)
    // fmt.Println("Service:", insurance.description, "Price:", insurance.monthlyFee * float64(insurance.durationMonths))
    
    // var e1 Expense = &Product{ name: "Kayak"}
    // var e2 Expense = &Product{ name: "Kayak"}
    // var e3 Expense = Service{ description: "Boat Cover"}
    // var e4 Expense = Service{ description: "Boat Cover"}
    //
    // fmt.Println("e1 == e2:", e1 == e2)
    // fmt.Println("e3 == e4:", e3 == e4)

    expenses := []Expense {
        Service{"Boat Cover", 12, 89.50, []string{}},
        Service{"Paddle Protect", 12, 8, []string{}},
        &Product{"Kayak", "Watersports", 275},
        Foo{"wangyao", 123},
    }

    for _, expense := range expenses {
        // s := expense.(Service)
        // fmt.Println("Service:", s.description, "Price:", s.monthlyFee * float64(s.durationMonths))
        if s, ok := expense.(Service); ok {
            fmt.Println("Service:", s.description, "Price:", s.monthlyFee * float64(s.durationMonths))
        } else {
            fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
        }

    }

    // dynamic type
    for _, expense := range expenses {
        switch s := expense.(type) {
            case Service:
                fmt.Println("Service:", s.description, "Price:", s.monthlyFee * float64(s.durationMonths))
            case *Product:
                fmt.Println("Product:", s.name, "Price:", s.price)
            default:
                fmt.Println("Expense:", s.getName(), "Cost:", s.getCost(true))
        }
    } 

}
