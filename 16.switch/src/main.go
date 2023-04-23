package main

import "fmt"

func main() {
	modulo := 6 % 2

	switch modulo {
	case 0:
		fmt.Println("Is pair")
	default:
		fmt.Println("Is odd")
	}

	switch modulo2 := 5 % 2; modulo2 {
	case 0:
		fmt.Println("Is pair")
	default:
		fmt.Println("Is odd")
	}

	// switch without argument

	value := 50

	switch {
	case value > 100:
		fmt.Println("is greater than 100")
	case value < 0:
		fmt.Println("is less than 0")
	default:
		fmt.Println("No condition")
	}

}
