package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// && means AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println(true)
	}

	// || means Or
	if valor1 == 1 || valor2 == 3 {
		fmt.Println(true)
	}

	// transform text to num
	value, err := strconv.Atoi("53")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(value)
	}
}
