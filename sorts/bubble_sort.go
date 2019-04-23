package main

import "fmt"

func BubbleSort(inputArray []int){
	lenInputArray:= len(inputArray)
	swap := true
	for ; swap == true; {
		swap = false
		for counter := 0; counter <  (lenInputArray - 1); counter ++ {
			elm1 := inputArray[counter]
			elm2 := inputArray[counter + 1]
			if elm1 > elm2 {
				swap = true
				inputArray[counter] = elm2
				inputArray[counter + 1] = elm1
			}
		}
		lenInputArray -= 1
	}
	fmt.Println(inputArray)
}

/*func main(){

	inputArray := []int{135305, 153880, 111255, 21406, 57002, 52054, 103056, 64209, 41616,
		175464, 147703, 72000, 190719, 13351, 137167, 88991, 186086, 99989, 131110, 99074, 8863,
		189535, 80037, 33947, 193112, 405, 188704, 22651, 32, 50654, 194172, 18810, 93988, 36912,
		1680, 152359, 118135, 55270, 13615, 110924, 147293, 75063, 58826, 51052, 137055, 80180,
		93078, 158963, 74302, 172829, 146623, 51337, 146759, 56251, 158533, 122770, 187570, 32842,
		74468, 130003, 67283, 41611, 40731, 2799, 130132, 96478, 46367, 68822, 24241, 76862, 7429,
		24395, 16003, 138646, 188248, 120352, 172003, 178577, 50200, 154099, 19591, 183367, 166508,
		83851, 176424}

	BubbleSort(inputArray)

	}*/