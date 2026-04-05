package main

import "fmt"

func main() {

	some := addition(12, 45)

	One, _ := multipleType()

	myClose := closer()

	fmt.Println("===========", myClose())
	fmt.Println("===========", myClose())

	fmt.Println(One)

	fmt.Println(some)

}

func addition(a, b int) int {
	return a + b
}

func multipleType() (int, string) {
	return 4, "yes i lke go!"
}

func closer() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
