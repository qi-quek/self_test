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
	//*strconv.Atoi is pretty much the same as strcon.Parseint
	//*except that for strcon.Parseint, you can set the base and int size
	//*etc base 10 and int 64 or int32
	input := scanner.Text()

	fmt.Println("you typed:", input)
}
