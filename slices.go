package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{2, 22, 4, 6}
	c := make([]int, 4)
	copy(c, a)

	b := make([]string, 3, 5)
	fmt.Println(a == nil, len(a), cap(a))
	b[1] = "cocuc"
	b[2] = "doudou"
	fmt.Println(b)
	fmt.Println("nnnnnnn", c)

	l := a[2:]
	fmt.Println("xxxxxx", l, slices.Equal(a, c))

	for o, z := range a {
		fmt.Println("=====))", z, "index:", o)
	}

}
