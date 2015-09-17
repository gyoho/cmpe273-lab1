package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

func(c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func(c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	length, width float64
}

func(r Rectangle) area() float64 {
	return r.length * r.width
}

func(r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	c := Circle{5}
	r := Rectangle{6, 4}
	shaperArr := [] Shaper{c, r}
	
	for _, elem := range shaperArr {
		fmt.Println("Shape details: ", elem)
		fmt.Println("Area: ", elem.area())
		fmt.Println("Perimeter: ", elem.perimeter())
	}
}

