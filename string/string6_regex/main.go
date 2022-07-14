package main

import (
	"fmt"
	"regexp"
)

func main() {
	//^ is like checking for things which are not, ^[]
	match, _ := regexp.MatchString("[^a-zA-Z_]", "peach")
	fmt.Println(match)
	if match == false {
		fmt.Println("does not contain any unwanted value")
	} else {

	}

	//*-----check for symbol-----------------------------
	match1, err := regexp.MatchString("[^a-zA-Z_]", "peach&*")

	fmt.Println(err)

	fmt.Println(match1)
	if match1 == false {
		fmt.Println("does not contain any unwanted value")
	} else {
		fmt.Println("string contains unwanted expression")
	}

	//*-----check for number-----------------------------
	match2, _ := regexp.MatchString("[^a-zA-Z_]", "peac123")

	fmt.Println(match2)
	if match2 == false {
		fmt.Println("does not contain any unwanted value")
	} else {
		fmt.Println("string contains unwanted expression")
	}

	//*-----check for space-----------------------------
	match3, _ := regexp.MatchString("[^a-zA-Z ]", "peac asd")

	fmt.Println(match3)
	if !match3 {
		fmt.Println("does not contain any unwanted value")
	} else {
		fmt.Println("string contains unwanted expression")
	}

}

//git test for CLI
