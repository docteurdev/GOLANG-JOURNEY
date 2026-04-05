package main

import (
	"fmt"
)

func main() {
	// 1- loop
	// i := 0
	// for i < len(myString) {
	// 	i++
	// 	fmt.Println("========", i)
	// }

	// for i := 0; i < len(myString); i++ {
	// 	fmt.Println("========", i)
	// }

	// for n := range len(myString) {

	// 	fmt.Printf("========q %d %c\n", n, myString[n])
	// }

	mathBest := bestStudent(19)
	fmt.Println("Math Best Student is:", mathBest)

}
func bestStudent(moy int) string {

	switch {
	case moy < 10:
		return "redoublant"
	case moy > 10 && moy <= 12:
		return "passable"
	case moy > 12 && moy <= 14:
		return "bien"
	default:
		return "excellent"
	}
}
