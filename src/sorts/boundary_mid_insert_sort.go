package main

import "fmt"

// BoundaryMidInsertSort - sorts int
func BoundaryMidInsertSort(inputArray []int) {
	lenInputArray := len(inputArray)
	var sortedArray []int
	sortedArray = append(sortedArray, inputArray[0])

	for counter := 1; counter < (lenInputArray); counter++ {
		stopMid := false
		lenSortedArray := len(sortedArray)
		currentElement := inputArray[counter]
		//fmt.Println(currentElement)

		if currentElement < sortedArray[0] {
			sortedArray = insertElement(sortedArray, currentElement, 0)
			continue
		}
		if currentElement > sortedArray[lenSortedArray-1] {
			sortedArray = insertElement(sortedArray, currentElement, lenSortedArray)
			continue
		}

		sortedSubarray := sortedArray
		midVal := len(sortedSubarray) / 2
		cursor := 0
		insertPosition := 0

		for stopMid == false {
			insertPosition = 0
			lenSortedSubarray := len(sortedSubarray) - 1

			if midVal == 0 {
				if currentElement > sortedSubarray[midVal] {
					insertPosition = cursor + 1
				} else {
					insertPosition = cursor
				}
				stopMid = true
				break
			} else if currentElement == sortedSubarray[midVal-1] {
				insertPosition = cursor + (midVal - 1)
				stopMid = true
				break
			} else if currentElement > sortedSubarray[midVal-1] {
				sortedSubarray = sortedSubarray[(midVal):(lenSortedSubarray + 1)]
				cursor = cursor + (midVal)
			} else if currentElement < sortedSubarray[midVal-1] {
				sortedSubarray = sortedSubarray[0:midVal]
			}

			insertPosition, midVal, stopMid = findMid(currentElement, sortedSubarray)
			insertPosition = insertPosition + cursor

		}

		sortedArray = insertElement(sortedArray, currentElement, insertPosition)
	}

	fmt.Println(sortedArray)
}

func insertElement(sortedArray []int, element int, index int) []int {
	newSlice := append(sortedArray, element)
	copy(newSlice[(index+1):], newSlice[index:])
	newSlice[index] = element
	return newSlice
}

func findMid(currentElement int, sortedSubarray []int) (int, int, bool) {
	lenSortedSubarray := len(sortedSubarray) - 1
	if currentElement < sortedSubarray[0] {
		return 0, 0, true
	} else if currentElement > sortedSubarray[lenSortedSubarray] {
		return lenSortedSubarray, 0, true
	}
	return 0, (lenSortedSubarray) / 2, false
}

/*func main(){

	input := []int {135305, 153880, 111255, 21406, 57002, 52054, 103056, 64209, 41616,
		175464, 147703, 72000, 190719, 13351, 137167, 88991, 186086, 99989, 150200,
		154099, 19591, 183367, 166508, 83851, 176424}

	BoundaryMidInsertSort(input)
}*/
