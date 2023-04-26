package main

import (
	"fmt"
	"sync"
)

/* goroutines
add go before calling a function
*/

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	say("Hi", &wg)
	go say("Go!", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Goodbye!")

	//time.Sleep(time.Second * 1)
}
