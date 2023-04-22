package main

import "fmt"

func main() {
	// docs: https://pkg.go.dev/fmt

	helloMessage := "Hello"
	goMessage := "GO"

	//println
	fmt.Println(helloMessage, goMessage)

	/* printf
	%d for integers
	%s for strings
	%v for unknown
	*/
	fmt.Printf("%s %s!\n", helloMessage, goMessage)

	/* sprintf return a string
	 */
	message := fmt.Sprintf("%s %s!", helloMessage, goMessage)
	fmt.Println(message)

	// var type using %T
	fmt.Printf("message type: %T", message)
}
