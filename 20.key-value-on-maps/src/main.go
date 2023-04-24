package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["jose"] = 14
	ages["john"] = 20
	ages["mary"] = 17

	fmt.Println(ages)

	// not deterministic way to travel
	for i, v := range ages {
		fmt.Println(i, v)
	}

	// find value by key
	value := ages["jose"]
	fmt.Println(value)
	// zero value non set key-value
	value2 := ages["joseph"]
	fmt.Println(value2)

	// ok for found key-value
	value3, ok := ages["joseph"]

	fmt.Println(value3, ok)
}
