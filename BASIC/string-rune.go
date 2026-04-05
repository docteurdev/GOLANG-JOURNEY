package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const d = "yes"
	fmt.Println(len(d))
	for i, l := range d {

		fmt.Println(i, l)
	}

	fmt.Println(utf8.RuneCountInString(d))

	for idx, runeValue := range d {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)
}
