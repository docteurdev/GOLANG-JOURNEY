package main

import "fmt"

func main() {
	i := 1
	fmt.Println("before zeroVal:", i)
	zeroVal(i)
	fmt.Println("after zeroVal==========:", i)

	fmt.Println("before zeroPtr:", i)
	zeroPtr(&i)
	fmt.Println("after zeroPtr:", i)
	fmt.Println("after zeroPtr:", &i)

	// the following lines will cause compilation error
	// because we cannot take the address of a literal
	//
}

func zeroVal(ival int) {
	ival = 0
	fmt.Println("inside zeroVal:", ival)
}

func zeroPtr(iptr *int) {
	*iptr = 0
}
