package three_int_sum

import (
	"fmt"
	"sort"
)

func main() {
	numMap := []int{-1, 0, 1, 2, -1, 4}
	val := 0
	result := find3NoList(numMap, val)
	fmt.Println(result)
}

func find3NoList(numArr []int, ans int) [][]int {
	ansList := [][]int{}
	sort.Ints(numArr)
	for index, val := range numArr {
		remainder := ans - val
		for i := index + 2; i < len(numArr); i++ {
			if numArr[index+1]+numArr[i] != remainder {
				continue
			}
			tempList := []int{val, numArr[index+1], numArr[i]}
			ansList = append(ansList, tempList)
		}
	}
	return ansList
}
