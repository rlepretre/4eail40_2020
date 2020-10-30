package main

import "fmt"

// Implement types Rectangle, Circle and Shape

type Rectangle struct { //create object with properties
	Width  int32
	Length int32
}

type Circle struct {
	Radius int32
}

type Shape interface {
	String() string
}

func (build Rectangle) String() string {
	return fmt.Sprintf("Rectangle of width %d and length %d", build.Width, build.Length)
}

func (build Circle) String() string {
	return fmt.Sprintf("Circle of radius %d", build.Radius)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		// Expected output:
		// Square of width 2 and length 3
		// Circle of radius 5
	}
}
