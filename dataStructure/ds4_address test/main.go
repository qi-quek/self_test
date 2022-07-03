package main

import "fmt"

func main() {
	testVal1 := 5

	//*if we do assign *(&testVal1), we are not assinging the pointer, we are assigning the value striaght
	testVal2 := *(&testVal1)

	//*so we have to reference towards address
	testVal3 := &testVal1

	fmt.Println("testval1(check1) - ", testVal1)
	fmt.Println("testval1(check1) - ", testVal2)
	fmt.Println("testval1(check1) - ", *testVal3)
	fmt.Println("------------------------------")
	// fmt.Println("")

	testVal1 = 6

	fmt.Println("------------------------------")
	fmt.Println("testval1(check2) - ", testVal1)
	fmt.Println("testval1(check2) - ", testVal2)
	fmt.Println("testval1(check2) - ", *testVal3)
	//*then deref address value only when we are calling the value

}
