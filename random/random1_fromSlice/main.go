package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	array := []string{"Metal", "Glass", "Plastic", "Paper"}
	rand.Seed(time.Now().Unix()) //initialise seed with time value

	fmt.Println(array[rand.Intn(len(array))])

	fmt.Println(rand.Intn(4750) + 250)
	//cannot random from a range of number via functions
	//so we do it in such a awy of intn(max-min)+min
	//so the value acquired will always be in the range of min-max

}
