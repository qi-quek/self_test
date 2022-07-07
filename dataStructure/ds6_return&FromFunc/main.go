package main

import "fmt"

func createVAl() *int {

	u := 2
	fmt.Printf("Type of testVal is %T, value of test val is %v\n, the address is %p", u, u, &u)
	return &u
}

func main() {

	testVal := createVAl()
	fmt.Println("line break--")
	fmt.Printf("Type of testVal is %T, value of test val is %v, the address is %p\n", *testVal, *testVal, testVal)

	fmt.Println("line break")

	testVal2 := createVAl()
	fmt.Println("line break--")
	fmt.Printf("Type of testVal is %T, value of test val is %v, the address is %p", *testVal2, *testVal2, testVal2)

}
