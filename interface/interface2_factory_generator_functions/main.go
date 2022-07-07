package main

import "fmt"

type Person struct {
	name string
	age  int
}

//return value is of name - func(name string); of type person
//the return value func is a function that can be called
//with the argument of type string
//so the age gets input first, thats why newbaby is of a function
//cant find out which data gets assigned first
//ultimately the concern is that the struct gets filled with the values.
func NewPersonFactory(name string) func(age int) Person {
	return func(age int) Person {
		return Person{
			name: name,
			age:  age,
		}
	}
}

func main() {
	newBaby := NewPersonFactory("john")
	fmt.Printf("%T\n", newBaby)
	baby := newBaby(1)
	fmt.Println(baby)

}
