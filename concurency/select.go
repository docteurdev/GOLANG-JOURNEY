package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(10 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(20 * time.Second)
		c2 <- "two"
	}()

	select {
	case c2 := <-c2:
		fmt.Println("yesss", c2)
	case c3 := <-c1:
		fmt.Println("yesss", c3)
	default:
		fmt.Println("no channel is ready")

	}

}
