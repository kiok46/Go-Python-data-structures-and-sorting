package main

import "fmt"

/*
Ok, so this is some new feature :)

attach some methods to structures.
*/

type shapes struct{
	width, height int
}


func (shape *shapes)area() int{
	return shape.height * shape.width
}


func (shape *shapes) perimeter() int{
	return shape.height + shape.width
}


func main() {
    var shape shapes

    shape = shapes{2, 3}
    fmt.Println(shape)
    fmt.Println("Area of the shape is: ", shape.area())
    fmt.Println("Perimeter of the shape is: ", shape.perimeter())
}