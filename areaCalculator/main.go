package main

import (
	"fmt"
	"math"
)

// 1. Define the interface
type Shape interface {
	Area() float64
}

// 2. Implement Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 3. Implement Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 4. Function that works with any Shape
func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area() // polymorphism in action
	}
	return total
}

// 5. Main function
func main() {
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 7},
		Rectangle{Width: 3, Height: 4},
	}

	total := TotalArea(shapes)
	fmt.Println("Total Area:", total)
}
