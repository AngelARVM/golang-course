package main

import "fmt"

type pc struct {
	ram int
	disk int
	brand string
}

func (myPc pc) ping(){
	fmt.Println(myPc.brand)
}

func (myPc *pc) duplicateRam(){
	myPc.ram*=2
}

//template
func main() {
	a := 50
	// & is for accessing memory address
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	// * is for accesing the value stored in memory address
	fmt.Println(*b)

	// when assign values using *, also change value stored in memory address
	*b = 100

	fmt.Println(*b)
	fmt.Println(a)

	myPc := pc{ram:16, disk: 240, brand: "macos"}

	fmt.Println(myPc)
	myPc.ping()

	myPc.duplicateRam()
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
}
