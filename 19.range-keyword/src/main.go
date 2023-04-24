package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) {
	var reverseText string

	for i := len(text) - 1; i >= 0; i-- {
		reverseText += string(text[i])
	}

	if strings.ToLower(reverseText) == strings.ToLower(text) {
		fmt.Println("its a palindrome")
	} else {
		fmt.Println("its not a palindrome")
	}
}

func main() {
	slice := []string{"hello", "go", "!"}

	for _, valor := range slice {
		fmt.Println(valor)
	}
}
