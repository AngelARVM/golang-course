package main

import "fmt"

func main() {
	// Defer. execute at the end of the app
	defer fmt.Println("Hola")
	fmt.Println("go!")

	// Continue & break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("Is 2")
			continue
		}

		if i == 6 {
			break
		}
	}
}
