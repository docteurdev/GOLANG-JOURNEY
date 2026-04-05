package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(10 * time.Second)

	<-timer1.C
	fmt.Println("here is the result of timer after 10 seconds: yes me and you")

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

}
