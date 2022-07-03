package main

import (
	"fmt"
)

// func fibonacci(ch chan int, quit chan bool) {
// 	x, y := 0, 1
// 	for {
// 		select {
// 		case ch <- x: //write to channel ch
// 			x, y = y, x+y
// 		case <-quit:
// 			fmt.Println("quit")
// 			return

// 		}
// 	}
// }

var flattened []int

func main() {
	flattened := []int{}
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}

	for _, v := range a {
		for _, k := range v {
			flattened = append(flattened, k)
		}
	}

	fmt.Println(flattened)
}

// func recursiveType(test123 interface{}) {

// 	if reflect.TypeOf(test123).Kind() != reflect.Int {
// 		for _, value := range test123 {
// 			recursiveType(value)
// 		}

// 	}

// 	for _, value := range test234 {
// 		if reflect.TypeOf(value).Kind() != reflect.Int {
// 			for _, value := range test234 {
// 				recursiveType(value)
// 			}
// 		} else {
// 			flattened = append(flattened, value)
// 		}
// 	}
// }

// {5,5,6,7,5,1,2,3,10}
// {5,{5,{6,7}}, {5,1,{2,3,{10}}}}

// index 0 = 5,{5,{6,7}}

// reflect.typeof=
