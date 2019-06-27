package main

import "fmt"


func mergeSort(inputArray []int){
	newInputArray := make([][]int, len(inputArray))
	var newMergedArray [][]int

	for index, elem := range inputArray {
		newInputArray[index] = append(newInputArray[index],elem)
	}

	for ; len(newInputArray) > 1; {
		//fmt.Println("in while", newInputArray)
		newMergedArray = nil
		for counter := 0; counter < (len(newInputArray) - 1); counter += 2 {
			subArray1 := newInputArray[counter]
			subArray2 := newInputArray[counter + 1]
			subMergedArray := mergeArraysAlternate(subArray1, subArray2)
			newMergedArray = append(newMergedArray, subMergedArray)
		}

		if (len(newInputArray) % 2) == 1{
			newMergedArray = append(newMergedArray, newInputArray[len(newInputArray) - 1])
		}

		//fmt.Println("merged :" , newMergedArray)
		newInputArray = newMergedArray
		//fmt.Println("new input :" , newInputArray)

	}

	fmt.Println(newInputArray[0])

}

func mergeArrays(subArray1 []int, subArray2[]int) (mergedArray []int){
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

func mergeArraysAlternate(subArray1 []int, subArray2[]int) (mergedArray []int){
	var newMergedArray []int
	sub1i, sub2i := 0,0

	for  ; (sub1i < len(subArray1) ) && (sub2i < len(subArray2)); {
		if subArray1[sub1i] <= subArray2[sub2i]{
			newMergedArray = append(newMergedArray, subArray1[sub1i])
			sub1i++
			continue
		}
		if subArray1[sub1i] > subArray2[sub2i]{
			newMergedArray = append(newMergedArray, subArray2[sub2i])
			sub2i++
		}
	}

	if sub1i < len(subArray1) {
		newMergedArray = append(newMergedArray, subArray1[sub1i:]...)
	}
	if sub2i < len(subArray2) {
		newMergedArray = append(newMergedArray, subArray2[sub2i:]...)
	}

	return newMergedArray
}

/*
func main(){

	input := []int{135305, 153880, 111255, 21406, 57002, 52054, 103056, 64209,
		41616, 175464, 147703, 72000, 190719}

	mergeSort(input)
}

*/
