package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string, 2)
	// ch := make(chan string, 10)

	//multiple values in a channel
	// buffered := make (chan string, 10)
	fmt.Printf("First initial Check %T, %v, %d\n", ch, ch, ch)
	fmt.Println("----------------")

	for i := 5; i < 20; i++ {
		go func(i int) {

			// the 2nd for loop is just to jumble up the value, no diff.
			// for j := 0; j < 10; j++ {
			// 	ch <- "Goroutine : " + strconv.Itoa(i)
			// }
			fmt.Println("check for pre channel", i)

			ch <- "Goroutine : " + strconv.Itoa(i)
			fmt.Printf("print1: %T, %v, %d\n", ch, ch, ch)

		}(i)

	}

	//*because goroutine is a seperate thread from main.go thread
	//*hence the for loop can continue
	//*once the channels are cleared,
	//*the other go rountines can continue
	//*no waitgroup, once the for loop has cleared 10 iterations of channel release
	//*main routine will end, which will end the all other routines

	//*this channel usage has no wait group, because we are using channel to
	//*determine release of the channel, hence wait group is not needed.

	// for k := 1; k <= 100; k++ {
	for k := 1; k <= 10; k++ {
		fmt.Println("check for print before channel")
		fmt.Println(k, <-ch)
	}

}
