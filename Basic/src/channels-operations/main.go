package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}
func main() {
	c := make(chan string, 2)
	c <- "msg1"
	c <- "msg2"
	fmt.Println(len(c), cap(c))
	// keywords Range and close
	//close
	close(c)
	//range
	for message := range c {
		fmt.Println(message)
	}
	//select
	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensage1", email1)
	go message("message2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("email recibido desde ", m1)
		case m2 := <-email2:
			fmt.Println("email recibido desde ", m2)
		}
	}
	close(email1)
	close(email2)
}
