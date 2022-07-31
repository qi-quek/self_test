package main

import (
	"fmt"
	"sort"
)

type user struct {
	name    string
	surname string
}

func main() {

	//declare and initilize the map with values
	users := map[string]user{
		"Roy":     {"Roy", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Jackson", "Michael"},
	}

	for key, value := range users {
		fmt.Println(key, value)
	}

	//to run the keys in order
	var keys []string
	for key := range users {
		keys = append(keys, key)
	}

	//sort the strings alphabetically
	sort.Strings(keys)

	//Walk through the keys and pull each value from the mp
	for _, key := range keys {
		fmt.Println(key, users[key])
	}

}
