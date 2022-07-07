package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	strings := []string{"hello", "world", "there"}

	wg.Add(len(strings))
	fmt.Println(len(strings))
	for _, val := range strings {
		//by inputting that we are looking for string argument string, the go-rountine will not skip the first 2 strings
		go func(string) {
			fmt.Println(val)
			wg.Done()
		}(val)

	}

	wg.Wait()
	fmt.Println("end of function")

}

//This prints only (1)"there", (2)"there", (3)"there"
//func main() {
//	var wg sync.WaitGroup
//	messages := []string{"hello", "world", "greetings"}
//	wg.Add(len(messages))
//	for _, msg := range messages {
//		go func() {
//			fmt.Println(msg)
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//}
