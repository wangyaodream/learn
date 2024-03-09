package main

import (
	"math"
	"math/rand"
)

func main() {
    val1 := 279.00
    val2 := 48.95

    Printfln("Abs: %v", math.Abs(val1))
    Printfln("Ceil: %v", math.Ceil(val2))
    Printfln("Copysign: %v", math.Copysign(val1, -5))
    Printfln("Floor: %v", math.Floor(val2))
    Printfln("Max: %v", math.Max(val1, val2))
    Printfln("Min: %v", math.Min(val1, val2))
    Printfln("Mod: %v", math.Mod(val1, val2))
    Printfln("Pow: %v", math.Pow(val1, val2))
    Printfln("Round: %v", math.Round(val1))
    Printfln("RoundToEven: %v", math.RoundToEven(val2))

    for i := 0; i < 5; i++ {
        Printfln("Value %v: %v", i,rand.Int())
    }

        
}
