package main

import "fmt"

func main() {
	fmt.Println(Proceeding)
}

type paymentProcessor int

const (
	Proceeding paymentProcessor = iota
	Success
	Error
)

func (py paymentProcessor) String() string {
	switch py {
	case Proceeding:
		return "processing"
	case Success:
		return "success"
	default:
		return " error"
	}
}
