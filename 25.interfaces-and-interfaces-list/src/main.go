package main

import "fmt"

type figure interface {
	area() float64
}

type Square struct {
	base float64
}

type Rectangle struct {
	base   float64
	height float64
}

func (c Square) area() float64 {
	return c.base * c.base
}

func (r Rectangle) area() float64 {
	return r.base * r.height
}

func calculate(f figure) {
	fmt.Println("Area:", f.area())
}

// template
func main() {
	mySquare := Square{base: 2}
	myRectangle := Rectangle{base: 2, height: 3}

	calculate(mySquare)
	calculate(myRectangle)

	myInterface := []interface{}{"hi", 12, 4.9}

	fmt.Println(myInterface...)
}
