package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// input := "foo   bar      baz"
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	fmt.Println("test1")
	// Array := [2]string{}

	scanner.Split(bufio.ScanWords)

	//it will forever be scanning
	for scanner.Scan() {
		// Array = append(Array, scanner.Text())
		fmt.Println(scanner.Text())
	}
	// fmt.Println(scanner.Text())
	// fmt.Println("zzz", Array)
	// fmt.Println("zzz", Array[0])
	// fmt.Println("zzz", Array[1])
	// // scanner.Scan()
}
