package main

import (
	"fmt"
)

func swapLexOrder(str string, pairs [][]int) string {

	if len(str) <= 1 || len(pairs) < 1 {
		return str
	}

	uniqueConnections := make(map[int]int)
	reversedStr := make([]byte, len(str))

	for pairIndex := 0; pairIndex < len(pairs); pairIndex++ {
		for start := 0; start < len(pairs[pairIndex]); start++ {
			base := pairs[pairIndex][start]
			for iterator := pairIndex + 1; iterator < len(pairs); iterator++ {

				if pairs[iterator][0] == base && pairs[iterator][1] == base {
					pairs = append(pairs[0 : iterator], pairs[iterator + 1 : ]...)
					fmt.Println(pairs)
					continue
				}

				if pairs[iterator][0] == base {
					_, ok := uniqueConnections[pairs[iterator][1]]
					if ok == true {
						pairs = append(pairs[0 : iterator], pairs[iterator + 1 : ]...)
						iterator--
						continue
					}
					pairs[pairIndex] = append(pairs[pairIndex], pairs[iterator][1])
					uniqueConnections[pairs[iterator][1]] += 1
					//fmt.Println("at first", base, pairs[iterator], pairs[pairIndex])
					pairs = append(pairs[0 : iterator], pairs[iterator + 1 : ]...)
					iterator--
					continue
				}
				if pairs[iterator][1] == base {
					_, ok := uniqueConnections[pairs[iterator][0]]
					if ok == true {
						pairs = append(pairs[0 : iterator], pairs[iterator + 1 : ]...)
						iterator--
						continue
					}
					pairs[pairIndex] = append(pairs[pairIndex], pairs[iterator][0])
					uniqueConnections[pairs[iterator][0]] += 1
					//fmt.Println("at second", base, pairs[iterator], pairs[pairIndex])
					pairs = append(pairs[0 : iterator], pairs[iterator + 1 : ]...)
					iterator--
					continue
				}
			}
			uniqueConnections[base] += 1
		}

		var toReverse []int
		for _, char := range pairs[pairIndex] {
			toReverse = append(toReverse, int(str[char - 1]))
		}
		BubbleSort(toReverse, "DES")
		BubbleSort(pairs[pairIndex], "ASC")
		for orderIndex, order := range pairs[pairIndex] {
			reversedStr[order - 1] = byte(toReverse[orderIndex])
		}

	}

	fmt.Println("to reverse", pairs)

	for i, chr := range reversedStr {
		if chr == 0 {
			reversedStr[i] = str[i]
		}
	}

	return string(reversedStr)
}

func BubbleSort(inputArray []int, order string){
	lenInputArray:= len(inputArray)
	swap := true
	for ; swap == true; {
		swap = false
		for counter := 0; counter <  (lenInputArray - 1); counter ++ {
			elm1 := inputArray[counter]
			elm2 := inputArray[counter + 1]
			if elm1 < elm2 && order == "DES"{
				swap = true
				inputArray[counter] = elm2
				inputArray[counter + 1] = elm1
			}
			if elm1 > elm2 && order == "ASC"{
				swap = true
				inputArray[counter] = elm2
				inputArray[counter + 1] = elm1
			}
		}
		lenInputArray -= 1
	}
	//fmt.Println(inputArray)
}

func permutation(inputLen int) [][]int {

	var permutationInput []int

	for i := 0; i < inputLen; i++ {
		permutationInput = append(permutationInput, i + 1)
	}

	permutations := make([][]int, 1)
	permutations[0] = append(permutations[0], permutationInput[inputLen - 1])

	for counter := len(permutationInput) - 2; counter >= 0; counter -- {

		appender := permutationInput[counter]

		for _, _ = range permutations {
			listItem := permutations[0]
			permutations = append(permutations[0 : 0], permutations[1:]...)

			for insertIndex := 0; insertIndex <= len(listItem); insertIndex ++ {
				tempList := make([]int, len(listItem) + 1)
				copy(tempList, listItem)
				tempList = append(tempList[0 : insertIndex + 1], tempList[insertIndex : len(tempList) - 1]...)
				tempList[insertIndex] = appender
				permutations = append(permutations, tempList)
			}

		}

	}

	return permutations
}

func recursive(input int) int {
	//fmt.Println(input)
	if input == 0 {
		return 0
	}
	return recursive(input - 1)
}

func findProfession(level int, pos int) string {
	// 0 - based
	level--
	pos--


	overallChange := false
	for level >= 0 {
		change := false
		if pos % 2 == 1 {
			change = true
		}
		//change := pos % 2 == 1
		fmt.Println("unknown", change)
		pos = pos / 2
		if change {
			overallChange = !overallChange
		}
		fmt.Println(overallChange)
		level--
	}
	if overallChange {
		return "Doctor"
	} else {
		return "Engineer"
	}
}

func main() {

	//		fzxmybhtuigowbyefkvhyameoamqei
	//str :=  "qvspxdrbvwfuaahtzbpjudfyzccgzwynkgihwmdshvfnvyvfjc"
	/*pairs := [][]int{{16,26},
	{2,25},
	{25,27},
	{19,20},
	{13,20},
	{4,26},
	{19,27},
	{18,26},
	{13,23},
	{1,4},
	{11,19},
	{16,19},
	{25,28},
	{19,30},
	{19,25},
	{1,11},
	{2,20},
	{10,22},
	{6,19},
	{7,26},
	{3,30},
	{15,23},
	{12,26},
	{1,3},
	{3,12},
	{3,26},
	{16,30},
	{2,16},
	{4,13}}
*/
	//fmt.Println(swapLexOrder(str, pairs))

	//xx := recursive(10)
	//fmt.Println(xx)

	fmt.Println(findProfession(10, 123))
}
