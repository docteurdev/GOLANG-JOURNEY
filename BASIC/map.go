package main

import (
	"fmt"
)

func main() {

	myMap := make(map[int]string)
	myMap[2] = "oumar"
	myMap[3] = "Alya"

	lovers := map[int]string{1: "my lovely alya", 2: "she heart Omar"}

	fmt.Println(myMap, lovers)
	for x, a := range lovers {
		fmt.Printf("%s ->", x, a)
	}

}
