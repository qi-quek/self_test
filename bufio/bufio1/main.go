package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var str string
	inp, _ := reader.ReadString('\n')
	fmt.Scanf("%s", &str)

	fmt.Println("z", inp)
	fmt.Println("zz", str)
}
