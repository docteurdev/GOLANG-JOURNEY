package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)
	p(now.Day())
	p(now.Year())

	p(now.Clock())
}
