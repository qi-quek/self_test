package main

import "fmt"

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	var result string
	stringEX := "Low Zhan Peng"

	for _, v := range stringEX {
		result = string(v) + result //this adds each iteration of "new-alphabet" to the start of the string
	}
	fmt.Println(result)

}
