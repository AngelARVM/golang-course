package main

import "fmt"

func main() {
	// array is defined with its length
	var array [4]int
	array[0] = 1
	array[1] = 2

	// len return the elements count in the array
	// cap return the max elements count for the array

	fmt.Println(array, len(array), cap(array))

	//slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(slice, len(slice), cap(slice))

	// Metodos
	fmt.Println(slice[0])   // index 0
	fmt.Println(slice[:3])  // index x < 3
	fmt.Println(slice[2:4]) // index 2 > x < 4
	fmt.Println(slice[4:])  // index 4 > x

	// Append
	slice = append(slice, 7) // adds second argument to the array in the first argument

	fmt.Println(slice)

	// append new list

	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)

	fmt.Println(slice)
}
