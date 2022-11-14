package main

import "fmt"

type tank interface {
	area() (string, float64)
	volume() float64
}

type value struct {
	radius float64
	height float64
	name   string
}

func (m value) area() (string, float64) {
	return m.name, 2*m.radius*m.height + 2*3.14*m.radius*m.radius
}

func (m value) volume() float64 {
	return 3.14 * m.radius * m.radius * m.height
}

func main() {

	var t tank
	t = value{10, 14, "hello"}
	areaName, area := t.area()
	fmt.Println("Area of tank ", areaName, area)
	fmt.Println("Volume of tank:", t.volume())

}
