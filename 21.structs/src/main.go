package main

import "fmt"

//structs
/*
objects doesnt exist in go, but there are structs that are a powerful way to define
*/

type Car struct {
	brand string
	year  int
}

func main() {
	myCar := Car{brand: "Ford", year: 2020}

	fmt.Println(myCar)

	var otherCar Car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
