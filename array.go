package main

import "fmt"

func main() {

	a := [5]int{2, 4, 6, 8, 9}

	b := [...]int{9, 8: 300, 3}
	fmt.Println("=========", len(a))

	fmt.Println("=========", b)

	// two dimensional array

	var d [2][3]int

	for l := range 2 {
		for v := range 3 {
			d[l][v] = l + v
		}
	}
	fmt.Println("cccc", d)
}
