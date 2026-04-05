package main

import "fmt"

func main() {

	m := 6
	for m < 10 {
		m++
		fmt.Println(m)
	}

	for x := 0; x < 8; x++ {
		fmt.Println("zzzzzzz", x)

	}
	for c := range 5 {
		fmt.Println("0000000000", c)
	}
}
