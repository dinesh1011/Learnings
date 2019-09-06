package main

import "fmt"

func mergeSortRecursive(inputArray []int) []int {

	if len(inputArray) <= 1{
		return inputArray
	}
	leftArray := mergeSortRecursive(inputArray[0 : (len(inputArray) / 2)])
	rightArray := mergeSortRecursive(inputArray[(len(inputArray) / 2) : ])

	mergedArray := mergeArrays1(leftArray, rightArray)
	return mergedArray

}

func mergeArrays1(subArray1 []int, subArray2[]int) (mergedArray []int){
	var newMergedArray []int

	for ; (len(subArray1) > 0 ) && (len(subArray2) > 0 ); {
		if subArray1[0] <= subArray2[0]{
			newMergedArray = append(newMergedArray, subArray1[0])
			subArray1 = subArray1[1:]
			continue
		}
		if subArray1[0] > subArray2[0]{
			newMergedArray = append(newMergedArray, subArray2[0])
			subArray2 = subArray2[1:]
		}
	}

	if len(subArray1) > 0 {
		newMergedArray = append(newMergedArray, subArray1...)
	}
	if len(subArray2) > 0 {
		newMergedArray = append(newMergedArray, subArray2...)
	}

	return newMergedArray
}


func main(){

	input := []int {135305, 153880, 111255, 21406, 57002, 52054}

	fmt.Println(mergeSortRecursive(input))
	fmt.Println("its from recursive merge sort")
	
}