package main

import "fmt"

type Rectangle struct {
	width, length float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func (r *Rectangle) scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main() {

	rect := Rectangle{width: 10, length: 5}
	fmt.Printf("Area of rectangle is %v\n", rect.area())

	rect.scale(2)
	fmt.Printf("Area of rectangle is %v\n", rect.area())

}
