package main

import "fmt"

var testArray []int
var mapCount map[int]int
var highestCount int

func init() {
	testArray = []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 3, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 9, 9, 9}
}

func main() {
	mapCount = make(map[int]int)
	for _, v := range testArray { //range through the array values

		if val, ok := mapCount[v]; !ok { //check if the value is already a key of the map
			//if the key is not present in the map

			mapCount[v] = 1 //assign the value as key to map, assign value as 1 - for first value

		} else { // for else statement means == ok

			//so assign value of map[v] to a
			a := val

			//assign back value of a + 1(counter increment by 1)
			mapCount[v] = a + 1 //assign back to map[v]
		}

	}

	// var highestCount int

	for key, _ := range mapCount {
		if highestCount == 0 {
			highestCount = mapCount[key]
		}

		if highestCount < mapCount[key] {
			highestCount = mapCount[key]
			continue
		}

	}

	sliceNum := []int{}

	for key, _ := range mapCount {
		if mapCount[key] == highestCount {
			sliceNum = append(sliceNum, key)
		}
	}

	fmt.Println("number", sliceNum, "has the highest count at", highestCount)

}
