// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sendMessage(num int, msgChan chan<- string) {

// 	fmt.Printf("Sending message %d ...\n", num)

// 	time.Sleep(time.Second * time.Duration(num))

// 	msg := fmt.Sprintf("Message %d is sent!", num)

// 	msgChan <- msg

// }

// func main() {

// 	sentMessage := make(chan string)

// 	go sendMessage(2, sentMessage)
// 	go sendMessage(4, sentMessage)
// 	go sendMessage(3, sentMessage)

// 	msg := <-sentMessage
// 	msg1 := <-sentMessage
// 	msg2 := <-sentMessage

// 	fmt.Println(msg, msg1, msg2)

// }

// ===================================
package main

import (
	"fmt"
	"time"
)

func sendEmail(num int, orderEmail chan<- string) {

	fmt.Printf(".... the email %d is getting send\n", num)

	time.Sleep(time.Second * time.Duration(num))

	email := fmt.Sprintf("the email %d is sent successfuly", num)

	orderEmail <- email

}

func main() {

	mailChannel := make(chan string)

	go sendEmail(1, mailChannel)
	go sendEmail(2, mailChannel)
	go sendEmail(3, mailChannel)

	mail1 := <-mailChannel
	mail2 := <-mailChannel
	mail3 := <-mailChannel

	fmt.Println(mail1)
	fmt.Println(mail2)
	fmt.Println(mail3)
	close(mailChannel)
}
