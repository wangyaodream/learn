package main

import (
    "fmt"
)

func calcWithTax(price float64) float64 {
    return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
    return price
}

func selectCalculator(price float64) func(float64) float64 {
    if (price > 100) {
        return calcWithTax
    }
    return calcWithoutTax

}

func main() {
    products := map[string]float64 {
        "Kayak": 275,
        "Lifejacket": 48.95,
    }

    for product, price := range products {
        var calcFunc func(float64) float64
        if (price > 100) {
            calcFunc = calcWithTax
        } else {
            calcFunc = calcWithoutTax
        }

        totalPrice := calcFunc(price)

        fmt.Println("Product:", product, "Price:", totalPrice)
    }
}
