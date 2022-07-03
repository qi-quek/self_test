package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	var numSlice []int

	wg.Add(1) //adding to waitgroup
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done() //done for this waitgroup add
	}()

	wg.Add(1) //adding to waitgroup
	go func() {
		for j := 0; j < 10; j++ {
			c <- j
		}
		wg.Done() //done for this waitgroup add
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for num := range c {
		numSlice = append(numSlice, num)
	}

	fmt.Println("hi there", numSlice)
}
