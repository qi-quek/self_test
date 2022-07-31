package main

import "fmt"

type player struct {
	name  string
	score int
}

func main() {
	players := map[string]player{
		"anna":  {"Anna", 42},
		"jacob": {"Jacob", 52},
	}

	player := players["anna"]
	player.score++
	players["anna"] = player
	fmt.Println(players)

}
