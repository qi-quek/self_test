package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	c := make(chan []string)
	var tempSlice []string
	scanner := bufio.NewScanner(os.Stdin)
	//scan the values of what we type and store it inside scanner
	fmt.Println("Enter category")

	// scanner.Scan()
	// var word string

	// fmt.Scanln(&word)
	// tempSlice = append(tempSlice, word)

	scanner.Scan()
	tempSlice = append(tempSlice, scanner.Text())

	// store the scanned line into variable input
	//*tempSlice = append(tempSlice, scanner.Text())
	fmt.Println("Enter weight")
	// scanner.Scan()

	// fmt.Scanln(&word)
	// tempSlice = append(tempSlice, word)

	scanner.Scan()
	tempSlice = append(tempSlice, scanner.Text())

	//Use channel as concurrency
	go func() {
		c <- tempSlice
		// c <- input2

		close(c) //close a channel after for loop has ended

	}()

	var output []string

	//*--
	for v := range c { //iterating through a channel gets the value from the channel and clears it
		//cannot assign to individual variables too
		output = v
	}
	fmt.Println("test word 1", output[0])
	fmt.Println("test word 2", output[1])
}
