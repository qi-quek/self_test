package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	array := []string{"tony", "yap", "left", "right"}
	var arrayNew []string
	// arrayNew2 := []string{}

	for _, val := range array {
		//to change all to lowercases
		//then uppercase the first letter
		newVal := cases.Title(language.Und, cases.NoLower).String(strings.ToLower(val))
		arrayNew = append(arrayNew, newVal)
	}
	fmt.Println("Printing first letter caps slice")
	fmt.Println(arrayNew)

	for index, _ := range array {
		array[index] = array[index][1:] + array[index][:1]
		// arrayNew[index] = arrayNew[index][1:] + arrayNew[index][:1]

	}
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Slice with all 1st letter pushed to each word's last letter")
	fmt.Println(array)
	// fmt.Println(arrayNew2)
}
