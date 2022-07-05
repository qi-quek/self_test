package main

import (
	"fmt"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	var result string
	stringEX := "Low Zhan Peng"

	for firstValCheck := range stringEX[:3] { //if only the first value is needed, etc. only index is needed and not the val
		fmt.Println("this is check for first value, without 2nd value", firstValCheck)
	}

	fmt.Println("clear space")

	//for _, v := range stringEX {
	//	result = string(v) + result //this adds each iteration of "new-alphabet" to the start of the string
	//}

	fmt.Println(result)

	for pos, char := range "日本\x80語" { //\x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

}
