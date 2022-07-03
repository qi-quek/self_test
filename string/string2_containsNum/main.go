package main

import (
	"fmt"
	"strings"
)

func main() {

	arrayNo := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	string := "hi there 1 there"
	string1 := "hi there there"
	string2 := "hi there5  there"

	for _, val := range arrayNo {
		if strings.Contains(string, val) {
			fmt.Println("string contains number")
		}

	}

	for _, val := range arrayNo {
		if strings.Contains(string1, val) {
			fmt.Println("string1 contains number")
		}
	}

	for _, val := range arrayNo {
		if strings.Contains(string2, val) {
			fmt.Println("string2 contains number")
		}
	}

}
