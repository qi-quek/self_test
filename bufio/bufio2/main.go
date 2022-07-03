package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//from the buf io package
	//make a new scanner object
	//that is an os.stdin
	//stdin means input
	scanner := bufio.NewScanner(os.Stdin)

	//scan the values of what we type and store it inside scanner
	scanner.Scan()

	//store the scanned line into variable input
	input := scanner.Text()

	// fmt.Println(input)
	fmt.Println("you typed:", input)
}
