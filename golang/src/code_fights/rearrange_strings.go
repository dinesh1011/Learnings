package main

import (
	"fmt"
)

func permutations(inputList []string) [][]string {

	inputLen := len(inputList) - 2
	outputList := make([][]string, 1)
	outputList[0] = append(outputList[0], inputList[inputLen+1])

	for counter := inputLen; counter >= 0; counter-- {
		appender := inputList[counter]
		for range outputList {
			baseArray := make([]string, len(outputList[0]))
			copy(baseArray, outputList[0])
			outputList = append(outputList[0:0], outputList[1:]...)

			for index := len(baseArray); index >= 0; index-- {

				baseArrayCopy := insert(baseArray, index, appender)
				outputList = append(outputList, baseArrayCopy)
				continue
			}
		}
	}

	return outputList
}

func insert(array []string, index int, element string) []string {
	array = append(array, "")
	array = append(array[0:index+1], array[index:len(array)-1]...)
	array[index] = element
	return array
}

func stringsRearrangement(inputArray []string) bool {

	var refArray []string
	refArray = append(refArray, inputArray[0])
	inputArray = append(inputArray[0:0], inputArray[1:]...)

	for count := 0; count < len(inputArray); {
		updated := false
		for index := len(refArray) - 1; index >= 0; index-- {

			if index == len(refArray)-1 && checkDiff(inputArray[count], refArray[index]) == true {
				refArray = append(refArray, inputArray[count])
				inputArray = append(inputArray[0:count], inputArray[count+1:]...)
				count = 0
				updated = true
				break
			}
			if index < len(refArray)-1 &&
				checkDiff(inputArray[count], refArray[index]) == true &&
				checkDiff(inputArray[count], refArray[index+1]) == true {
				refArray = append(refArray, inputArray[count])
				inputArray = append(inputArray[0:count], inputArray[count+1:]...)
				count = 0
				updated = true
				break
			}
			if index == 0 && checkDiff(inputArray[count], refArray[index]) == true {
				refArray = append(refArray, "")
				refArray = append(refArray[0:1], refArray[0:len(refArray)-1]...)
				refArray[0] = inputArray[count]
				inputArray = append(inputArray[0:count], inputArray[count+1:]...)
				count = 0
				updated = true
				break
			}
		}

		if updated == false {
			count++
		}
	}

	fmt.Println("input ", inputArray)
	fmt.Println("ref ", refArray)
	if len(inputArray) > 0 {
		return false
	}
	return true
}

func checkDiff(str1 string, str2 string) bool {
	invalidCount := 0
	for index, char1 := range str2 {
		if string(char1) != string(str1[index]) {
			invalidCount++
		}
	}
	if invalidCount == 1 {
		return true
	}
	return false
}

/*
func main() {

	inputArray := []string{"abc",
		"xbc",
		"axc",
		"abx"}
	//xx := permutations(inputArray)

	//fmt.Println(xx)
	//fmt.Println(len(permutations(inputArray)))

	fmt.Println(stringsRearrangement(inputArray))

}

*/
