package main

import (
	"fmt"

	pk "golang-course/22.access-modifier/src/myPackage"
)

/*
can set as public or private
if the name of func or struct is capital then its public, else its private
*/

func main() {
  var myCar pk.CarPublic
  myCar.Brand = "Ferrari"
  fmt.Println(myCar)
}
