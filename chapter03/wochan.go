package main

import (
	"fmt"
	"sync"
)

func createCashier(cashierID int, wg *sync.WaitGroup) func(int) {
	orderProcessed := 0
	return func(orderNum int) {
		if orderProcessed < 10 {
			// Cashier is ready to serve!
			//fmt.Println("Cashier ", cashierID, "Processing order", orderNum, "Orders Processed", orderProcessed)
			fmt.Println(cashierID, "->", orderProcessed)
			orderProcessed++
		} else {
			fmt.Println("Cashier ", cashierID, "I am tired! I want to take rest!", orderNum)
		}
		wg.Done()
	}
}

func main() {
	cashierIndex := 0
	var wg sync.WaitGroup

	// make cashier 3
	cashiers := []func(int){}
	for i := 1; i <= 3; i++ {
		cashiers = append(cashiers, createCashier(i, &wg))
	}

	for i := 0; i < 30; i++ {
		wg.Add(1)

		cashierIndex = cashierIndex % 3
		func(cashiers func(int), i int) {
			go cashiers(i)
		}(cashiers[cashierIndex], i)

		cashierIndex++
	}

	wg.Wait()
}
