package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	// queue := make(chan string)
	queue <- "one"
	queue <- "two"
	close(queue)

	fmt.Printf("Type - %T\nvalue - %v\nD - %d\n", queue, queue, queue)
	fmt.Println("size of queue:", len(queue))

	for elem := range queue { //*range because it is a buffered channel of 2
		// fmt.Println("what the fuck is wrong with windows", elem)
		fmt.Println(elem)
	}

	//*A deadlock happens when a group of goroutines
	//*are waiting for each other and none of them is able to proceed
}
