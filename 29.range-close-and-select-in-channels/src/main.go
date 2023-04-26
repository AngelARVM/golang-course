package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

//template
func main() {
	c := make(chan string, 2)
	c <- "message1"
	c <- "message2"

	fmt.Println(len(c), cap(c))

	// Close the channel so cant add more data

	close(c)
	// c <- "message2"

	// Range travels the channel
	for message := range c {
		fmt.Println(message)
	}

	// Select. Used when channels responses order are unknown
	email1 := make(chan string)
	email2 := make(chan string)

	go message("email1", email1)
	go message("email2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println(m1)
		case m2 := <-email2:
			fmt.Println(m2)
		}
	}

}
