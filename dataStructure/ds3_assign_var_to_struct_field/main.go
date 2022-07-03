package main

import "fmt"

func main() {
	type test struct {
		test *bool
	}

	var valBool bool

	// var testMap = make(map[string]bool)

	// testMap["valid"] = false

	//*get the value and assigning
	newStruct := test{(&valBool)}
	// newStruct := test{testMap["valid"]}

	fmt.Println("struct check 1 -", newStruct.test) //*tested and confirmed, this is assigning addres to field
	fmt.Println("struct check 1 -", *(newStruct.test))
	fmt.Println("bool check a-", valBool)
	// fmt.Println("bool check i-", testMap["valid"])

	*(&valBool) = true
	// testMap["valid"] = true

	fmt.Println("struct check 2 -", newStruct.test) //*tested and confirmed, this is assigning addres to field
	fmt.Println("struct check 2 -", *(newStruct.test))
	fmt.Println("bool check b-", valBool)
	// fmt.Println("bool check ii-", testMap["valid"])

}
