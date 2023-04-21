package main

import "fmt"

func main() {
	// Linea de comentario

	// Constants are defined with reserved word "const" followed by type and its assignment
	const pi float64 = 3.141592
	const pi2 = 3.15

	// Variables may be assigned without declaring by using :=
	base := 12

	// Also can be defined with reserved word "var", just as const
	var height int = 14

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)
	fmt.Println("base:", base)
	fmt.Println("height:", height)

	// Go wont compile if unused variables are declared

	// Zero values: variables non assigned has a default value
	var a int     // 0
	var b float64 // 0
	var c string  // ""
	var d bool    // false

	fmt.Println(a, b, c, d)

	// calculate square area
	var area int = base * height

	// fmt.PrintF and % symbol are used to formatted printing
	fmt.Printf("A square with base %d and heigh %d has an area of %d ", base, height, area)

}
