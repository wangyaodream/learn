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

func generateSymbolArray(symbols map[string]uint) []string {
    symbolArr := []string{}
    for symbol, count := range symbols {
        for i := uint(0); i < count; i++ {
            symbolArr = append(symbolArr, symbol) 
        }
    }
    return symbolArr
}

func getSpin(reel )

func main() {
    symbols := map[string]uint{
        "A": 4,
        "B": 7,
        "C": 12,
        "D": 20,
    }

    symbolArr := generateSymbolArray(symbols)

    fmt.Println(symbolArr)


    balance := uint(200)

    for balance > 0 {
        bet := getBet(balance)
        if bet == 0 {
            break
        }
        balance -= bet
    }

    fmt.Println("Bye!")

}
