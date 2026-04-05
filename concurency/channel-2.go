package main

import (
	"fmt"
	"sync"
)

type TransRequest struct {
	id      string
	balance int
}

func debitMoneyFromAccount(trans TransRequest, debit int, debitC chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("is the user %d exist?????", trans.id)
	fmt.Printf("is the user has enough balance....? \n ")
	fmt.Printf("the money: %d is debited from the user account successfully 🌽\n", trans.balance)

	debitC <- debit
}
func creditMoneyToAccount(trans TransRequest, credit int, creditC chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("money is getting crediting----- %s \n", trans.id)
	fmt.Printf("the money: %d is credited to the user account successfully 🥕\n", trans.balance)
	creditC <- credit
}

func main() {
	wg := sync.WaitGroup{}
	debitChan := make(chan int)
	creditC := make(chan int)
	transAmont := 200
	transaction := TransRequest{id: "tr-1", balance: 50000}

	wg.Add(1)
	go debitMoneyFromAccount(transaction, transAmont, debitChan, &wg)

	debitedAmont := <-debitChan

	fmt.Println("-------------------------------------\n \n")

	wg.Add(1)
	go creditMoneyToAccount(transaction, transAmont, creditC, &wg)

	creditAmount := <-creditC

	wg.Wait()

	fmt.Printf("the transaction %s got debited with the amount %d and got credted of %d", transaction.id, debitedAmont, creditAmount)

}
