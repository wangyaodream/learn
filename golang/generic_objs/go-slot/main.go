package main

import "fmt"



func getBet(balance uint) uint {
    var bet uint
    for {
        fmt.Printf("Enter your bet, or 0 to quit (balance = $%d):", balance)
        fmt.Scan(&bet)

        if (bet > balance) {
            fmt.Println("Bet cannot be larger than balance.")
        } else {
            break
        }
    }
    return bet
}

func main() {
    balance := 200
    bet := getBet(uint(balance))
    if bet != 0 {
        fmt.Println("The bat:", bet)
    } else {
        fmt.Println("Bye!")
    }
}
