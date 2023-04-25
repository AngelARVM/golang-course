package main

import "fmt"

type pc struct {
	ram int
	disk int
	brand string
}

func (myPc pc) String() string {
	return fmt.Sprintf("{\n  \"ram\": %d\n  \"disk\": %d\n  \"brand\": \"%s\"\n}", myPc.ram, myPc.disk, myPc.brand)
}


//template
func main() {
	myPc := pc{ram:16, disk: 240, brand: "macos"}

	fmt.Println(myPc)

}
