package main

import "fmt"

/* functions
defined with camel case
*/
func helloGo(a int) {
	fmt.Printf("Hello GO! %d\n", a)
}

func funcTest(a, b int, c string) {
	fmt.Printf("%d %d %s\n", a, b, c)
}

func sum(a, b int) int {
	return a + b
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	helloGo(1)
	helloGo(2)
	helloGo(3)

	funcTest(1, 2, "hello")

	fmt.Println(sum(3, 4))

	// multiple single line assignment
	value1, value2 := doubleReturn(2)

	fmt.Printf("value1:%d value2:%d\n", value1, value2)

	// take first value
	value3, _ := doubleReturn(3)

	fmt.Println(value3)
}
