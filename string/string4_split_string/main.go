package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	someString := "one    two   three four "

	words := strings.Fields(someString)

	fmt.Println(words, len(words)) // [one two three four] 4

	var newWordArr []string

	// words = strings.ReplaceAll(words," ",",")
	for _, word := range words {
		fmt.Println(word)
		newWordArr = append(newWordArr, cases.Title(language.Und, cases.NoLower).String(strings.ToLower(word)))
	}

	fmt.Println(newWordArr)
	fmt.Printf("%T, %v of string files\n", newWordArr, newWordArr)

	stringFiles := strings.Join(newWordArr[:], " ")
	fmt.Printf("%T, %v  of stringFiles\n", stringFiles, stringFiles)

	wordTest := []string{"hi", "test", "123"}

	fmt.Println(wordTest)
}
