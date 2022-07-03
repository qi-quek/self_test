package main

func main() {

}

func mergeSort(refSlice []int) []int { //recursive function - calling itself
	if len(refSlice) < 2 {
		return refSlice
	} // if there is only 1 item in the slice, return back the slice value, no need for further merging

	//spitting the input slice into 2 halves, this is the top half, with index ranging from first data to middle data
	topHalf := mergeSort(refSlice[:len(refSlice)/2])
	//*

	//spitting the input slice into 2 halves, this is the bottom half, with index ranging from where input Slice was cut off from firstHalf
	//to end of the slice
	bottomHalf := mergeSort(refSlice[len(refSlice)/2:])
	//*

	return merge(topHalf, bottomHalf)
}

func merge(topHalf []int, bottomHalf []int) []int {
	var tempMergeSlice []int                                               //variable for temp slice
	var topHalfIndex int                                                   //int value for top half index
	var bottomHalfIndex int                                                //int value for bottom half index
	for topHalfIndex < len(topHalf) && bottomHalfIndex < len(bottomHalf) { //keeps looping thruogh if index cycle value is still less than the reference slice
		if topHalf[topHalfIndex] < bottomHalf[bottomHalfIndex] { //if value of bottom half is bigger than top half
			tempMergeSlice = append(tempMergeSlice, topHalf[topHalfIndex]) //append smaller value first - so that it is in ascending order
			topHalfIndex++                                                 //++ index count for cycle through

		} else { //if bottom half value is smaller than tophalf value
			tempMergeSlice = append(tempMergeSlice, bottomHalf[bottomHalfIndex]) //append smaller value first - so that it is in ascending order
			bottomHalfIndex++                                                    //++ index count for cycle through
		}
	}
	for ; topHalfIndex < len(topHalf); topHalfIndex++ { //this condition is for when the merge sort split is uneven, where there are base- 3 numbers instead of 2
		tempMergeSlice = append(tempMergeSlice, topHalf[topHalfIndex])
	}
	for ; bottomHalfIndex < len(bottomHalf); bottomHalfIndex++ { //this condition is for when the merge sort split is uneven, where there are base- 3 numbers instead of 2
		tempMergeSlice = append(tempMergeSlice, bottomHalf[bottomHalfIndex])
	}
	return tempMergeSlice
}
