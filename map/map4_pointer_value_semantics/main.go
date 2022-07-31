package main

import "fmt"

// type players struct{
// 	name string
// 	score int
// }

func main() {
	scores := map[string]int{
		"anna":  21,
		"jacob": 12,
	}
	double(scores, "anna")

	fmt.Println("Score:", scores["anna"])

}

//*we move map around with value semantics
//*
func double(scores map[string]int, player string) { //double gets it own copy of the map.
	scores[player] = scores[player] * 2 //even though we are using value semnatics
	//*any sort of reading and writing for maps, uses pointer semantics.
}
