package main

import "fmt"

func main() {
	//for
	fmt.Println("For")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// for while
	counter := 0
	fmt.Println("For While")
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	counterForever := 0
	fmt.Println("For Forever")
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever == 10 {
			break
		}
	}
}
