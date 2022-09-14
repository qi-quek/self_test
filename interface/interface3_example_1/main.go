package main

import "fmt"

// https://www.geeksforgeeks.org/interfaces-in-golang/

// Interface is a custom type that is used to specify a set of one or more method signatures and the interface is abstract,
//You are not allowed to create an instance of the interface.

//intance means like,
//type hellos struct{
//	abc string
//	def int
//}
//
//a := &hellos{
//	abc: "hello",
//	def: 20
//}

//test

//but you can however, createa a variable of an interface type and be assigned to a concrete type value that has the methods the interface requires

type tank interface {
	Tarea() float64  //return value of the func
	Volume() float64 //return value of the func
}

type myvalue struct {
	radius float64
	height float64
}

func (m myvalue) Tarea() float64 {
	return 2*m.radius*m.height + 2*3.14*m.radius*m.radius
}

func (m myvalue) Volume() float64 {
	return 3.14 * m.radius * m.radius * m.height
}

func main() {
	var t tank //create a variable of the interface type
	t = myvalue{10, 14}
	fmt.Println("Area of tank:", t.Tarea())
	fmt.Println("Volume of tank:", t.Volume())

}
