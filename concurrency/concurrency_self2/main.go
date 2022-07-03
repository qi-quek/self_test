package main

import "fmt"

func generator() <-chan int { //returning a channel that can only send an int
	c := make(chan int) //unbuffered channel
	//*buffered channel has values eg; c := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ { //channel is closed once 10 values have been sent into the channel, 0 to 9
			c <- i
		}
		close(c) //close a channel after for loop has ended

		//main thread should end once the first value into the channel clears

	}()

	return c
}

func receiver(c <-chan int) { //*a sending channel
	fmt.Println("length of c is ", len(c))

	//for i:=0; i<=1; i++
	//*for loop of a channel, receives values from the channel repeatedly until it is closed
	for v := range c { //iterating through a channel gets the value from the channel and clears it
		fmt.Println(v)
	}
	//only when the channel is closed, then will it terminate
	//c := [1,1,1,1,1,1,1]
	//c := [0]
}

func main() {
	c := generator()

	receiver(c) //0->1->2->3->..9

	//0<-c

	//1)we either wg.wait(), or we use a seperate for loop to wait

	//line 34 ends slower than input of line 32?

}

//A deadlock happens when a group of goroutines
//are waiting for each other and none of them is able to proceed

//---------Explanation-------
//if channel is not closed - channel is still opened will cause deadlock
//it is waiting to receive a value, then receiver is also waiting for a value, both routines
//are waiting for a value, hence deadlock
