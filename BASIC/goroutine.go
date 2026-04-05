package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	tableNumber int
	prepTime    time.Duration
}

func processOrder(order Order) {
	fmt.Printf("Preparing order for table d%...\n", order.tableNumber)

	time.Sleep(order.prepTime)

	fmt.Printf("order ready  for table d%!\n\n", order.tableNumber)
}

func main() {

	orders := []Order{
		{tableNumber: 1, prepTime: 2 * time.Second},
		{tableNumber: 2, prepTime: 3 * time.Second},
		{tableNumber: 3, prepTime: 1 * time.Second},
		{tableNumber: 4, prepTime: 2 * time.Second},
		{tableNumber: 5, prepTime: 4 * time.Second},
	}

	wg := sync.WaitGroup{}

	for _, order := range orders {

		wg.Add(1)

		go func() {
			defer wg.Done()
			processOrder(order)
		}()
	}

	wg.Wait()
	fmt.Println("All task are done")
}
