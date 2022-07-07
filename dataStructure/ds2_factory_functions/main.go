package main

import "fmt"

//*Manual creation of Person struct
type Person1 struct {
	Name string
	Age  int
}

//method of the person struct
func (p Person1) Greet() {
	fmt.Printf("Hi! My Name is %s\n", p.Name)
}

func main() {
	//instantiate a struct
	Q := Person1{"Hello", 30}

	//Method call
	Q.Greet()

	fmt.Println("end of normal struct call")
	fmt.Println("Start of factory functions")

	L := NewPerson("larry", 30)
	L.Greet1()
	fmt.Printf("%T", L)

}

func (p person) Greet1() {
	fmt.Printf("Hi! My name is %s", p.name)
}

type person struct {
	name string
	age  int
}

//Newperson function
/*to assign values to a new variable
useful when you need all fields of the struct to be filled
and dont want to worry about missing any field instantiation
*/
func NewPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}
