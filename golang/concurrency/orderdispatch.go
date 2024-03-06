package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customer = []string{"Alice", "Bob", "Charlie", "Dora"}

func DispatchOrders(channel chan DispatchNotification) {
	source := rand.NewSource(time.Now().Unix())
	generator := rand.New(source)
	orderCount := rand.Intn(3) + 2
	fmt.Println("Order count:", orderCount)
	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification{
			Customer: Customer[generator.Intn(len(Customer)-1)],
			Product:  ProductList[generator.Intn(len(ProductList)-1)],
			Quantity: generator.Intn(10),
		}
	}
	close(channel)
}
