package main

import "fmt"

func main() {
	x := 1
	y := 3
	z := 10
	fmt.Printf("x=%d y=%d z=%d", x, y, z)
	fmt.Println()

	result := y + z
	fmt.Println("sum x+z", result)

	result = y - z
	fmt.Println("min y-z", result)

	result = y * z
	fmt.Println("times y*z", result)

	result = z / y
	fmt.Println("div z/y", result)

	result = z % y
	fmt.Println("mod z%y", result)

	x++
	fmt.Println("x++", x)
}
