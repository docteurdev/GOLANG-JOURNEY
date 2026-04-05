package main

import "fmt"

func main() {
	firstAccount := Account{"oumar", 1000}
	firstAccount.deposit(500)
	firstAccount.withdraw(200)
	fmt.Println(firstAccount.balance)
}

type Account struct {
	owner   string
	balance int
}

func (a *Account) deposit(amount int) {
	a.balance += amount
}

func (a *Account) withdraw(amount int) {
	a.balance -= amount
}
