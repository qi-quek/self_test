package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("boring!") // calling the boring function, and using the word "boring" as the string
	for i := 0; i < 5; i++ {
		fmt.Printf("you say: %q\n", <-c) //prints out the word from the channel
		//q is a quoted character
	}
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string) <-chan string { //output or return value is receive only channel of strings{
	c := make(chan string)
	go func() { //launching go routine from inside the function.
		for i := 0; ; i++ { //no upper limit as it is restricted by the main function of the sending of the channel
			c <- fmt.Sprintf("%s %d", msg, i) //sending a string with the message and counter into the channel
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

	}()
	return c //return the channel to the caller

}
