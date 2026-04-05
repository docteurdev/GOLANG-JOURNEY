package main

func main() {
	newBankAccount := BankAccount{"ajdi", 1000}
	newMobileMoney := MobileMoney{"coucou", 400}

	processPayment(&newBankAccount, 100)

	processPayment(&newMobileMoney, 300)

}

type Payment interface {
	pay(amount int)
}

type BankAccount struct {
	owner   string
	balance int
}

type MobileMoney struct {
	owner   string
	balance int
}

func processPayment(p Payment, amount int) {
	p.pay(amount)
}

func (b *BankAccount) pay(amount int) {
	b.balance -= amount
}

func (b *MobileMoney) pay(amount int) {
	b.balance -= amount
}
