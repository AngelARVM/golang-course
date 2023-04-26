package main

import "fmt"

/*
	Allows to share information between goroutines
*/

func say(text string, c chan<- string) {
	// <- means that data will be shared on the channel c
	c <- text
}

func main() {
	// type channel, zise
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c)
}
