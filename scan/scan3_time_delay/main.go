package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("hello")
	duration := time.Duration(3) * time.Second
	time.Sleep(duration)

	fmt.Println("there")
}
