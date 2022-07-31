package main

import "fmt"

func main() {
	var map_1 map[int]int

	map_2 := map[int]int{}

	var map_3 = make(map[int]int)

	fmt.Println("Value of map with only declaration", "is =>", map_1)
	fmt.Println("Value of map initialised with empty map, using :=", "is =>", map_2)
	fmt.Println("Value of map with make(map[int]string)", "is =>", map_3)

	map_2[5] = 10
	map_3[5] = 10

	//*the line below is uncommented as it prompts a panic
	//*because it is a nil map and cannot be assigned values
	//*uncomment the line below and see what happens
	// map_1[5] = 10

	fmt.Println("after assigning values")
	fmt.Println("Value of map with only declaration", "is =>", map_1)
	fmt.Println("Value of map initialised with empty map, using :=", "is =>", map_2)
	fmt.Println("Value of map with make(map[int]string)", "is =>", map_3)

}
